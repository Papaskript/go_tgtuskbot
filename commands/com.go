package commands

const (
	Start      = "/start"
	Add        = "/add"
	Remove     = "/remove"
	ShowNote   = "/show"
	UpdateNote = "/update"
)

var StartMessage = "Привет, здесь ты можешь созать заметку!\n" +
	"Нажми кнопку:" + " " + "Menu"

var AddMessage = "Отправьте вашу заметку"
var RemoveMessage = "Отправьте номер заметки, которую хотите удалить"
var UpdateMessage = "Отправьте номер заметки,которую треубуеться обновить"
var NewUpMessage = "Отправьте ваши изменения"
var SuccessfullyUpdated = "Заметка успешно обновлена!"
