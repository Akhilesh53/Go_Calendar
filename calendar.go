package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const dateFormat = "02/01/2006" //DD/MM/YYYY

func create_callback_string(action string,year string, month string, day string) string {
	if len(day) == 1{
		day = "0"+day
	}
	if len(month) == 1{
		month = "0"+month
	}
	return  action + " " + day + " " + month + " " + year 
}

func separateCallbackData(callback string) (string,string,string,string){
	str := strings.Split(callback, " ")
	action := str[0]
	day := str[1]
	month := str[2]
	year := str[3]

	return action,day,month,year
}

func getYear(time time.Time) int{
	return time.Year()
}

func getDay(time time.Time) int{
	return time.Day()
}

func getMonth(time time.Time) time.Month{
	return time.Month()
}

func getLocation(time time.Time) time.Location{
	return *time.Location()
}

func getMonthFromMonthNumber(monthNumber int) time.Month{
	return time.Month(monthNumber)
}

func createCalender(year int, month int)tgbotapi.InlineKeyboardMarkup{

    curr := time.Now()
	currentLocation := getLocation(curr)

	Month := getMonthFromMonthNumber(month)

	 firstDay := time.Date(year, Month,1, 0, 0, 0, 0, &currentLocation)
	 lastDay := firstDay.AddDate(0, 1, -1)

	 startDay := 0
	 _,_,endDay := lastDay.Date()

	 fmt.Println(startDay, endDay)

    var calender = tgbotapi.NewInlineKeyboardMarkup(createCalenderInterface(firstDay,lastDay,startDay,endDay)...)

	return calender
}

func createCalenderInterface(first time.Time,last time.Time,startDay, endDay int)[][]tgbotapi.InlineKeyboardButton{
	var tempRow = make([][]tgbotapi.InlineKeyboardButton,0)
	month := getMonth(first)
	year := strconv.Itoa(getYear(first))
	day := strconv.Itoa(getDay(first))
	var temp = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("<",create_callback_string("PREV_MONTH",year,strconv.Itoa(int(month)),day,)),
		tgbotapi.NewInlineKeyboardButtonData(month.String()[0:3]+" "+year,create_callback_string("IGNORE",year,strconv.Itoa(int(month)),day,)),
		tgbotapi.NewInlineKeyboardButtonData(">",create_callback_string("NEXT_MONTH",year,strconv.Itoa(int(month)),day,)),
	)
	tempRow = append(tempRow,temp)

	temp = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Sun",create_callback_string("IGNORE",year,strconv.Itoa(int(month)),day,)),
		tgbotapi.NewInlineKeyboardButtonData("Mon",create_callback_string("IGNORE",year,strconv.Itoa(int(month)),day,)),
		tgbotapi.NewInlineKeyboardButtonData("Tue",create_callback_string("IGNORE",year,strconv.Itoa(int(month)),day,)),
		tgbotapi.NewInlineKeyboardButtonData("Wed",create_callback_string("IGNORE",year,strconv.Itoa(int(month)),day,)),
		tgbotapi.NewInlineKeyboardButtonData("Thu",create_callback_string("IGNORE",year,strconv.Itoa(int(month)),day,)),
		tgbotapi.NewInlineKeyboardButtonData("Fri",create_callback_string("IGNORE",year,strconv.Itoa(int(month)),day,)),
		tgbotapi.NewInlineKeyboardButtonData("Sat",create_callback_string("IGNORE",year,strconv.Itoa(int(month)),day,)),

	)
	tempRow = append(tempRow,temp)

	var k = 0

    var row = tgbotapi.NewInlineKeyboardRow()
	for i:=startDay; i <= endDay; i++ {
			Date := first.AddDate(0,0,i)

			if checkSameMonth(first,Date){
				dayStr := getDayString(getDay(Date))
				
				for k < 7{
					if (tempRow[1][k]).Text == Date.Weekday().String()[0:3]{
						row = append(row,tgbotapi.NewInlineKeyboardButtonData(dayStr,create_callback_string("DAY",strconv.Itoa(getYear(Date)),strconv.Itoa(int(getMonth(Date))),strconv.Itoa(getDay(Date)),))) 
					    k++
						break 
					}else {
						row = append(row,tgbotapi.NewInlineKeyboardButtonData("  ",create_callback_string("IGNORE",strconv.Itoa(getYear(Date)),strconv.Itoa(int(getMonth(Date))),strconv.Itoa(getDay(Date)),)))
					    k++
					}
				}
			}else{
				row = append(row,tgbotapi.NewInlineKeyboardButtonData("  ",create_callback_string("IGNORE",strconv.Itoa(getYear(Date)),strconv.Itoa(int(getMonth(Date))),strconv.Itoa(getDay(Date)),)))
			}
		if k == 7{
			k = 0
			tempRow = append(tempRow, row)
			row = tgbotapi.NewInlineKeyboardRow()
		}
	}
	for k < 6{
		row = append(row,tgbotapi.NewInlineKeyboardButtonData(" ",create_callback_string("IGNORE",strconv.Itoa(getYear(first)),strconv.Itoa(int(getMonth(first))),strconv.Itoa(getDay(first)),)))
	    k++
	}
    tempRow = append(tempRow, row)
	return tempRow
}

func checkSameMonth(first, date time.Time) bool{
	return first.Month() == date.Month()
}

func getDayString(day int) string{
	if day < 10{
		return "0"+strconv.Itoa(day) 
	}
	return strconv.Itoa(day)
}
 
func CalendarAction(update tgbotapi.Update, Bot *tgbotapi.BotAPI){
	
	var bot = Bot
	if update.Message != nil { 
		if update.Message.Text == "/Calendar"{
			curr := time.Now()
			year := getYear(curr)
			month := getMonth(curr)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyMarkup = createCalender(year,int(month),)

			bot.Send(msg)
		}else{
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,"Invalid Command")
				bot.Send(msg)
		}
	}else if update.CallbackQuery != nil {  
		action,day,month,year := separateCallbackData(update.CallbackQuery.Data)
		
		switch action{
			case "IGNORE":
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID," ")
				if _, err := bot.Request(callback); err != nil {
					panic(err)
				}
			case "DAY":
				date  := day+"/"+month+"/"+year
				var t,_ = time.Parse(dateFormat, date)

				today := time.Now().Local()
	
				if t.Local().After(today) {				
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,"Choosen Date is Invalid")
					msg.ReplyMarkup = createCalender(getYear(time.Now()),int(getMonth(time.Now())),)
				    bot.Send(msg)
				}else if t.Local().Before(today) {
					if t.Weekday().String() == "Saturday" || t.Weekday().String() == "Sunday" {
						fmt.Println("Invalid "+ t.Weekday().String())
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,"Choosen Date is "+ t.Weekday().String())
				        msg.ReplyMarkup = createCalender(getYear(time.Now()),int(getMonth(time.Now())),)
						bot.Send(msg)
					}else{
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,"Choosen Date is "+ day + " " + month + " " + year)
						bot.Send(msg)
					}
				}
			case "PREV_MONTH":
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID,"Previous Month")
				if _, err := bot.Request(callback); err != nil {
					panic(err)
				}
				date := day+"/"+month+"/"+year
				Date,_ := time.Parse(dateFormat, date)

				prevDate := Date.AddDate(0,-1,0)
				prevMonth := int(getMonth(prevDate))
				prevYear := getYear(prevDate)
				msg := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID,update.CallbackQuery.Message.MessageID,createCalender(prevYear,prevMonth,))
				bot.Send(msg)
			
			case "NEXT_MONTH":
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID,"Next Month")
				if _, err := bot.Request(callback); err != nil {
					panic(err)
				}
				date := day+"/"+month+"/"+year
				Date, _ := time.Parse(dateFormat, date)
				prevDate := Date.AddDate(0,1,0)
				prevMonth := int(getMonth(prevDate))
				prevYear := getYear(prevDate)
				msg := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID,update.CallbackQuery.Message.MessageID,createCalender(prevYear,prevMonth,))
				bot.Send(msg)

			default:
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,"Invalid Command")
				bot.Send(msg)
		}		
	}
}

func main(){
    
	// CREATE A NEW INSTANCE OF THE BOT
	bot, err := tgbotapi.NewBotAPI("YOUR TELEGRAM BOT TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// GET UPDATES FROM THE TELEGRAM
	// ANY INCOMING MESSAGE WILL BE UPDATE HERE
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// WHATEVER THE ACTION SPECIFY, PERFORM ACTION ACCORDING TO THAT
		CalendarAction(update,bot)
	}
}