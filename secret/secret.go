package secret

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"github.com/wisaitas/share-pkg/stringutil"
)

// Create a separate viper instance for secrets
var secretViper = viper.New()

func init() {
	_, errCheckYAML := os.Stat(PATH_SECRET_YAML)
	_, errCheckYML := os.Stat(PATH_SECRET_YML)

	isExistYAML := true
	if os.IsNotExist(errCheckYAML) {
		isExistYAML = false
	}

	isExistYML := true
	if os.IsNotExist(errCheckYML) {
		isExistYML = false
	}

	if isExistYAML {
		readConfig(PATH_SECRET_YAML)
	}

	if isExistYML {
		readConfig(PATH_SECRET_YML)
	}

	settings := secretViper.AllSettings()
	fmt.Printf("settings: %+v\n", settings)
}

func readConfig(path string) error {
	secretViper.SetConfigFile(path)

	if err := secretViper.ReadInConfig(); err != nil {
		return fmt.Errorf("[Share Package Secret] : %w", err)
	}

	return nil
}

func ReadSecret(param any) error {
	if param == nil {
		return errors.New("[Share Package Secret] : val is nil")
	}

	val := reflect.ValueOf(param)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return errors.New("[Share Package Secret] : param must be a struct")
	}

	return processStruct(val, "", "")
}

func processStruct(val reflect.Value, viperPrefix string, envPrefix string) error {
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		tagValue := typ.Field(i).Tag.Get(TAG_SECRET)
		field := val.Field(i)

		snakeFieldName := stringutil.ToSnakeCase(fieldName)
		viperKey := snakeFieldName

		if viperPrefix != "" {
			viperKey = viperPrefix + "." + snakeFieldName
		}

		fieldEnvName := stringutil.ToScreamingSnakeCase(fieldName)
		envKey := fieldEnvName

		if envPrefix != "" {
			envKey = envPrefix + "_" + fieldEnvName
		}

		envValue := os.Getenv(envKey)

		valueFound := false

		if envValue != "" {
			if err := setFieldValue(field, envValue); err != nil {
				return fmt.Errorf("[Share Package Secret] : %w", err)
			}
			valueFound = true
		}

		if !valueFound && secretViper.IsSet(viperKey) {
			if err := setFieldValue(field, secretViper.GetString(viperKey)); err != nil {
				return fmt.Errorf("[Share Package Secret] : %w", err)
			}
			valueFound = true
		}

		if !valueFound && tagValue != "" {
			if err := setFieldValue(field, tagValue); err != nil {
				return fmt.Errorf("[Share Package Secret] : %w", err)
			}
		}

		if field.Kind() == reflect.Struct {
			newViperPrefix := viperKey
			newEnvPrefix := envKey

			if err := processStruct(field, newViperPrefix, newEnvPrefix); err != nil {
				return fmt.Errorf("[Share Package Secret] : %w", err)
			}
		}
	}

	return nil
}

func setFieldValue(field reflect.Value, value string) error {
	if !field.CanSet() {
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int:
		num, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("[Share Package Secret] : %w", err)
		}
		field.SetInt(int64(num))
	case reflect.Bool:
		if value == UPPER_TRUE || value == LOWER_TRUE || value == ONE {
			field.SetBool(true)
		} else if value == UPPER_FALSE || value == LOWER_FALSE || value == ZERO {
			field.SetBool(false)
		} else {
			return errors.New("[Share Package Secret] : invalid bool value")
		}
	case reflect.Int64:
		duration, err := time.ParseDuration(value)
		if err != nil {
			return fmt.Errorf("[Share Package Secret] : %w", err)
		}
		field.SetInt(int64(duration))
	}

	return nil
}
