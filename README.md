# Doc_Service

## Сборка проекта

### Импорт html ресурсов

```text
mklink /j "D:\projects\src\projects\doc\doc_service\web" "D:\projects\src\projects\doc\doc_service_gui\dist"
```

Отвечает за формирование документов по заранее составлены  шаблонам

## Добавление шаблона в программу

Шаблоны хранятся в папке doc_templates, она содержит подкаталоги  каждый из которых отвечает за свой формат шаблона, в ней лежит карта шаблонов с описанием, кодом и путем к шаблону.

Пример содержимого list.json

```json
{
    "type": "docx_list",
    "resource": {
        "1": {"file": "dynamic_table_tpl.docx", "comment": "Демонстрирует возможности формирования документов"}
    }
}
```

По аналогии описав новый файл его требуется положит рядом с файлом list.json.

## Запрос на заполнение файла docx

```go
//Установка соединения с сервисом
tr, err := transport.NewTransport("127.0.0.1:8030")
if err != nil {
    log.Println("создание транспорта", err)
}

//Запрос на генерацию шаблона 
res, err := tr.DocxPerform("1", methods.TParams{NameFile: "test", ConvertPDF: true, Rotation: true},
map[string]interface{}{
    "col_labels": []string{"fruit", "vegetable", "stone", "thing"},
    "tbl_contents": []interface{}{
        map[string]interface{}{"label": "yellow", "cols": []string{"banana", "capsicum", "pyrite", "taxi"}},
        map[string]interface{}{"label": "red", "cols": []string{"apple", "tomato", "cinnabar", "doubledecker"}},
        map[string]interface{}{"label": "green", "cols": []string{"guava", "cucumber", "aventurine", "card"}},
    },
})

if err != nil {
    log.Println("Отправка данных", err)
}
```

Функция tr.DocxPerform принимает ряд параметров:

- code - код шаблона информация по коду находиться в файле  list.json.
- params принимает структуру methods.TParams с описанием параметров обработки файлов
- data получает карту значений для шаблона в формате map[string]interface{}

Возвращает res *methods.TFile и тело ошибки

Описание структуры

```go
    type TFile struct {
        Path     string //Путь где храниться файл
        Ext      string `json:"ext"`      // Формат отчета
        Name     string `json:"name"`     // Название
        FileData []byte `json:"fileData"` // Данные файлы
    }
```

## Запрос с несколькими файлами

```go
tr, err := transport.NewTransport("127.0.0.1:8030")
if err != nil {
    log.Println("создание транспорта", err)
}

group, err := tr.SendGroupFile()
if err != nil {
    log.Println("инициализация открытия", err)
}

err = group.DocxPerform("1", methods.TParams{Rotation: false, NameFile: "claim", Join: true}, map[string]interface{}{
    "col_labels": []string{"fruit", "vegetable", "stone", "thing"},
    "tbl_contents": []interface{}{
        map[string]interface{}{"label": "yellow", "cols": []string{"banana", "capsicum", "pyrite", "taxi"}},
        map[string]interface{}{"label": "red", "cols": []string{"apple", "tomato", "cinnabar", "doubledecker"}},
        map[string]interface{}{"label": "green", "cols": []string{"guava", "cucumber", "aventurine", "card"}},
    },
})
if err != nil {
    log.Println("Добавление файла", err)
}

if err != nil {
    log.Println("Добавление файла", err)
}

err = group.XlsxPerform("1", methods.TParams{Rotation: true, NameFile: "claim_table", Join: true}, formats_xlsx.TValuesRender{
    Data: map[string]interface{}{
        "manager_name": "test",
    },
    DataTable: []map[string]interface{}{
        {
            "A": "5110204838/086943 от 30.06.2023",
            "B": "93 336,82",
        },
    },
})
if err != nil {
    log.Println("Добавление файла", err)
}

res, err := group.Send()
if err != nil {
    log.Println("Отправка данных", err)
}
```

После формирования списка вызываем функцию отправки, в ответ получим []*methods.TFile и тело ошибки.
