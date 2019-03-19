# r8calc

This is a *very* simple intial implementation of a tool do to some calculations regarding train power requirements for the [Run8 train simulator](http://www.run8studios.com/).  It needs a Run8 train save file (in XML) and a copy of [Deezle's train makeup spreadsheets](https://www.thedepotserver.com/forums/threads/bnsf-and-up-train-makeup-spreadsheets.4515/) from The Depot.

To run, run `r8calc` on the command line with the name of the XML file and the name of the spreadsheet as arguments, in that order.  For example, if you have saved Deezle's spreadsheet as "TrainCalculators.xlsx" and your train save file is "MyTrain.xml", your command would be `r8calc  MyTrain.xml TrainCalculators.xlsx`.

This will parse your train file, open the spreadsheet in Excel, and populate all the data fields for both BNFS and UP with data taken from your train file.  *Please note*:  For now, it is best if this, the spreadsheet, and the train file are all in the same directory.  If people find this useful I'll add more intelligent path handling in the future.

Thanks to:  The Depot, from whose [Rolling Stock Database](https://www.thedepotserver.com/reference/rollingstock/) I grabbed car empty weight information, and whose forum is a great resource for folks like me who are just learning this sim; and to Deezle for taking the time to put their great tool together.

You're free to do with this as you wish, but please keep credit to The Depot as long as the information they painstakingly gathered is in the `carWeight()` function.