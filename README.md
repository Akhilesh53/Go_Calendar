# Go Calendar

A Calendar Interface using Telegram Bot API in GoLang.

# SET UP

**REQUIREMENTS**
=======================================================

Get your Token from **BOTFATHER** on Telegram

**STEPS TO BE FOLLOWED FOR THIS PROJECT SETUP**
=======================================================

**STEP 1** : Create a new Folder in your PC

**STEP 2** : Clone the repository from the link below.
         If you dont have GO setup already in your PC
         Download GO

**STEP 3** : Clone the Repo from the Link

**STEP 4** : Run Go mod init - to setup workspace (necessary if you are creating your own files)

**STEP 5** : Run Go mod tidy

**STEP 6** : Run Go mod vendor - (run at last to have all the necessary libraries)

**STEP 7** : Write / Modify the Code

**STEP 8** : Create new branch with your name.

            git branch <your_name>
            git checkout <your_name>

**STEP 9** : Stage your changes and Commit
         
            git add <File.go>
            git commit -m "<your commit message>"

**STEP 10** : Push the code
             
             git.push
          
          
# FUNCTIONS

1) **create_callback_string(action string,year string, month string, day string) string**
   
   This function creates a callback string for inline buttons

   Params :
   
   action : Action that Calendar will perform
            a) IGNORE - Donot perform any action
            b) DAY - Date button (1-31) is pressed
            c) PREV_MONTH - Move to Previous month
            d) NEXT_MONTH - Move to Next month
   year : Value of year
   month : Value of moth
  
   Result : Combined string of all Params

2) **separateCallbackData(callback string) (string,string,string,string)**
   
   This function separates the callback string 

   Params : A callback string
   
   Result : Separate callback string in action,year,month,date

3) **func GetYear(time time.Time) int**
4) **func GetDay(time time.Time) int**
5) **func GetMonth(time time.Time) time.Month**
6) **func GetLocation(time time.Time) time.Location**

    These are utility functions to get the value of day, month, year andd location based on current time
    
    Param :
    time : time.Time value

    Result : Value accrrding to function

7) **getMonthFromMonthNumber(monthNumber int) time.Month**

    This function returns the Name of Month based on its integer value
    E.g : 7 = July, 8 = August

8) **CreateCalender(year int, month int)tgbotapi.InlineKeyboardMarkup**

    This functions generates a calender interface using inline buttons

    Params :
    year : value of year for which calender is be to created
    month : value of month for which calender is be to created
   
    Result : A calender Interface

11) **createCalenderInterface(stringToBeAppend string,first time.Time,last time.Time,startDay, endDay int)[][]tgbotapi.InlineKeyboardButton**
  
    A helper function of CreateCalender() to generate calendar Interface

12) **CalendarAction(update tgbotapi.Update, Bot * tgbotapi.BotAPI)**
    
    This function is to perform actions on the Calendar Interface

    Params : 
    update : Telegram update value
    Bot : TGBOT for sending messages
    
# WORKING

**NOTE:**   1) Saturday and Sunday are connsidered to be as invalid (Change accordingly).

            2) Any date selected after today is invalid.
            
            In both the the cases a new calender will be created for current month and year

Send **/Calendar** message with two params year and month to generate the calendar for that year and month

If the action specified is  IGNORE - Nothing will be performed.

                            DAY - A date is selected ( When we press a date, DAY Action is performed).
                            
                            PREV_MONTH - Generate calender for previous month ( When we press on **<**).
                            
                            NEXT_MONTH - Generate calender for next month (When we press on **>**).
                            

# RESULTS

https://user-images.githubusercontent.com/64018679/183618292-f2b7cc21-05d4-4d84-beb1-67b8c1eca9b4.mp4

https://user-images.githubusercontent.com/64018679/183618483-4f44325d-7018-40ec-b528-8dc8a27363ee.mp4

https://user-images.githubusercontent.com/64018679/183618546-2e0976af-15cd-4e23-b22b-585395278030.mp4

https://user-images.githubusercontent.com/64018679/183618600-b58aed0a-4051-48f6-ba80-77e205b4c11a.mp4

