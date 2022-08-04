package main

import (
	"regexp"
)

type DotnetAppSettingsRemover struct {
	AppSettings     map[string]interface{}
	EnvironmentData string
}

func (s DotnetAppSettingsRemover) RemoveEnvVariable() string {

	var removeEnvVariable func(accum string, m map[string]interface{}, envFilesContent string) string

	// Depth-First Search Keys defined in the AppSettings Json and replace the same variables in the EnvironmentFiles Content
	removeEnvVariable = func(accum string, m map[string]interface{}, envFilesContent string) string {

		var env_name string
		var res string = envFilesContent

		for k, v := range m {

			if len(accum) > 0 {
				env_name = accum + "__" + k
			} else {
				env_name = k
			}

			switch v.(type) {
			default:
				panic("json read error")
			case map[string]interface{}:
				res = removeEnvVariable(env_name, v.(map[string]interface{}), res)
			case interface{}:
				mc := regexp.MustCompile("(?m)[\r\n]+^" + env_name + "=.*$")
				res = mc.ReplaceAllString(res, "")
			}
		}

		return res
	}

	return removeEnvVariable("", s.AppSettings, s.EnvironmentData)
}
