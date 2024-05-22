package widget

import (
	"errors"
)

type Translations struct {
	Open            string
	Save            string
	ShowHiddenFiles string
	File            string
	Cancel          string
}

func getEnglishTranslations() Translations {
	return Translations{
		Open:            "Open",
		Save:            "Save",
		ShowHiddenFiles: "Show Hidden Files",
		File:            "File",
		Cancel:          "Cancel",
	}
}

func getRussianTranslations() Translations {
	return Translations{
		Open:            "Выбрать",
		Save:            "Сохранить",
		ShowHiddenFiles: "Показать скрытые файлы",
		File:            "Файл",
		Cancel:          "Отмена",
	}
}

func LoadNameLocalization(language string) (Translations, error) {
	if language == "EN" {
		return getEnglishTranslations(), nil
	} else if language == "RU" {
		return getRussianTranslations(), nil
	} else {
		return getEnglishTranslations(), ErrLanguageNotFound
	}

}

var ErrTranslationNotFound = errors.New("folder's translation hasn't found")
var ErrLanguageNotFound = errors.New("language hasn't found")

func getRusMapFolders() map[string]string {
	return map[string]string{
		"Documents": "Документы",
		"Downloads": "Загрузки",
		"Music":     "Музыка",
		"Pictures":  "Картинки",
		"Videos":    "Видео",
		"Users":     "Пользователи",
	}
}

func GetRusNameFolder(engNameFolder string) (string, error) {
	mapping := getRusMapFolders()
	if val, ok := mapping[engNameFolder]; ok {
		return val, nil
	}
	return engNameFolder, ErrTranslationNotFound
}
