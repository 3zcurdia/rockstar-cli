package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var (
	adjs = []string{
		"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark",
		"summer", "icy", "delicate", "quiet", "white", "cool", "spring", "winter",
		"patient", "twilight", "dawn", "crimson", "wispy", "weathered", "blue",
		"billowing", "broken", "cold", "damp", "falling", "frosty", "green",
		"long", "late", "lingering", "bold", "little", "morning", "muddy", "old",
		"red", "rough", "still", "small", "sparkling", "throbbing", "shy",
		"wandering", "withered", "wild", "black", "young", "holy", "solitary",
		"fragrant", "aged", "snowy", "proud", "floral", "restless", "divine",
		"polished", "ancient", "purple", "lively", "nameless",
	}
	nouns = []string{
		"waterfall", "river", "breeze", "moon", "rain", "wind", "sea", "morning",
		"snow", "lake", "sunset", "pine", "shadow", "leaf", "dawn", "glitter",
		"forest", "hill", "cloud", "meadow", "sun", "glade", "bird", "brook",
		"butterfly", "bush", "dew", "dust", "field", "fire", "flower", "firefly",
		"feather", "grass", "haze", "mountain", "night", "pond", "darkness",
		"snowflake", "silence", "sound", "sky", "shape", "surf", "thunder",
		"violet", "water", "wildflower", "wave", "water", "resonance", "sun",
		"wood", "dream", "cherry", "tree", "fog", "frost", "voice", "paper",
		"frog", "smoke", "star",
	}
	repos = []string{
		"proyect", "system", "sys", "app", "web", "site", "api",
		"cli", "framework", "extention", "ios", "android", "bot", "ui",
		"3D", "tracer", "manager", "admin", "orm", "db", "nosql", "saas",
		"example", "test",
	}
	messages = []string{
		"No time to commit.. My people need me!", "fixed mistaken bug", "mergederp", "Who has two thumbs and remembers the rudiments of his linear algebra courses?  Apparently, thisguy.",
		"Derp", "Committing fixes in the dark, seriously, who killed my power!?", "I did it for the lulz!", "bara bra grejjor", "Now added delete for real", "Fucking templates.", "MACIEJ, WE WENT OVER THIS. EXPANDTAB.", "I am sorry", "fixed the israeli-palestinian conflict", "pay no attention to the man behind the curtain", "hey, what's that over there?!", "Trying to fake a conflict", "hoo boy", "Don't push this commit", "/sigh", "hey, look over there!", "fix bug, for realz", "lolwhat?", "FOR REAL.", "making this thing actually usable.", "LOL!", "bla", "Handled a particular error.", "Fixed the fuck out of #946!", "??! what the ...", "LAST time, serjoscha87, /dev/urandom IS NOT a variable name generator...", "The last time I tried this the monkey didn't survive. Let's hope it works better this time.", "clarify further the brokenness of C++. why the fuck are we using C++?", "foo", "oops - thought I got that one.", "MOAR BIFURCATION", "fix /sigh", "I don't get paid enough for this shit.", "Final commit, ready for tagging", "oops!", "forgot we're not using a smart language", "Become a programmer, they said. It'll be fun, they said.", "Todo!!!", "Corrected mistakes", "Yep, Shubham was right on this one.", "if you're not using et, fuck off", "arrgghh... damn this thing for not working.", "Fixed a little bug...", "Well, it's doing something.", "Replace all whitespaces with tabs.", "typo", "A fix I believe, not like I tested or anything", "woa!! this one was really HARD!", "oops, forgot to add the file", "First Blood", "RAINER, WE WENT OVER THIS. EXPANDTAB.", "This really should not take 19 minutes to build.", "changes", "John made me do it", "giggle.", "extra debug for stuff module", "It worked for me...", "Yep, minimal was right on this one.", "it's friday", "I honestly wish I could remember what was going on here...", "Locating the required gigapixels to render...", "syntax", "SHIT ===> GOLD", "Fixing Lyntor's bug.", "rats", "I don't give a damn 'bout my reputation", "Another commit to keep my CAN streak going.", "And a commit that I don't know the reason of...", "Please no changes this time.", "assorted changes", "This branch is so dirty, even your mom can't clean it.", "LOTS of changes. period", "add actual words", "That last commit message about silly mistakes pales in comparision to this one", "It's 2015; why are we using ColdFusion?!", "hmmm", "unionfind is no longer being molested.", "I should get a raise for this.", "asdfasdfasdfasdfasdfasdfadsf", "Done, to whoever merges this, good luck.", "Derp. Fix missing constant post rename", "Is there an award for this?", "pointless limitation", "Spinning up the hamster...", "yo recipes", "fixed conflicts (LOL merge -s ours; push -f)", "#GrammarNazi", "SOAP is a piece of shit", "omg what have I done?", "tunning", "bifurcation", "Fix my stupidness", "Reticulating splines...", ":(:(", "some brief changes", "better ignores", "Your commit is writing checks your merge can't cash.", "DEAL WITH IT", "Updated", "things occurred", "Push poorly written test can down the road another ten years", "Added translation.", "8==========D", "Feed. You. Stuff. No time.", "someday I gonna kill someone for this shit...", "Derpy hooves", "Fucking submodule bull shit", "Future self, please forgive me and don't hit me with the baseball bat again!", "Does not work.", "A long time ago, in a galaxy far far away...", "This bunny should be killed.", "breathe, =, breathe", "lol digg", "commit", "lots of changes after a lot of time", "well crap.", "I really should've committed this when I finished it...", "happy monday _ bleh _", "de-misunderestimating", "Either Hot Shit or Total Bollocks", "changed things...", "include shit", "lol", "Herpy dooves.", "touched...", "Glue. Match sticks. Paper. Build script!", "My bad", "Does anyone read this? I'll be at the coffee shop accross the street.", "Merging the merge", "No changes made", "added security.", "more debug... who overwrote!", "totally more readable", "copy and paste is not a design pattern", "Pig", "debug line test", "It compiles! Ship it!", "It'd be nice if type errors caused the compiler to issue a type error", "did everything", "more ignored words", "I'm human", "Working on tests (haha)", "Switched off unit test 4 because the build had to go out now and there was no time to fix itproperly.",
		"remove debug", "all good", "bump to 0.0.3-dev:wq", "I'm too foo for this bar", "See last commit", "ffs", "oopsie B|", "another big bag of changes", "really ignore ignored worsd", "Fix the fixes", "[Insert your commit message here. Be sure to make it descriptive.]", ".", "arrrggghhhhh fixed!", "For the sake of my sanity, just ignore this...", "I know, I know, this is not how I’m supposed to do it, but I can't think of something better.", "this doesn't really make things faster, but I tried", "put code that worked where the code that didn't used to be", "move your body every every body", "herpderp", "work in progress", "Issue #5 is now Issue #26", "pgsql is more strict, increase the hackiness up to 11", "tagging release w.t.f.", "Added missing file in previous commit", "pep8 - cause I fell like doing a barrel roll", "It's Working!", "just checking if git is working properly...", "more stuff", "buenas those-things.", "Herping the fucking derp right here and now.", "Nitpicking about alphabetizing methods, minor OCD thing", "debugo", "apparently i did something…", "jobs... steve jobs", "Yep, Ali was right on this one.", "(c) Microsoft 1988", "WAHYU, WE WENT OVER THIS. C++ IO SUCKS.", "that's all folks", "Just stop reading these for a while, ok..", "This is supposed to crash", "epic", "I CAN HAZ COMMENTZ.", "fix", "I'm totally adding this to epic win. +300", "Too lazy to write descriptive message", "LUKASZ, WE WENT OVER THIS. C++ IO SUCKS.", "Shit code!", "Ali made me do it", "It's getting hard to keep up with the crap I've trashed", "TONY SUCKS", "derpherp", "Herpderp, shoulda check if it does really compile.", "Things went wrong...", "I should have had a V8 this morning.", "Herp derp I left the debug in there and forgot to reset errors.", "Yes, I was being sarcastic.", "This solves it.", "Gross hack because L1NT doesn't know how to code", "just shoot me", "debug suff", "Whatever.", "first blush", "Programming the flux capacitor", "fixed some minor stuff, might need some additional work.", "TODO: write meaningful commit message", "I cannot believe that it took this long to write a test for this.", "Trust me, it's not badly written. It's just way above your head.", "Not one conflict, today was a good day.", "Made it to compile...", "I am even stupider than I thought", "magic, have no clue but it works", "Commit committed....", "Finished fondling.", "No changes after this point.", "Last time I said it works? I was kidding.  Try this.", "Does this work", "WTF is this.", "GIT :/", "Obligatory placeholder commit message", "Something fixed", "Herping the derp", "gave up and used tables.", "...", "Switched off unit test 8 because the build had to go out now and there was no time to fix itproperly.",
		"Yep, makerbot was right on this one.", "more ignores", "I can't believe it took so long to fix this.", "this should fix it", "FIX", "Whee.", "IEize", "Added a banner to the default admin page. Please have mercy on me =(", "I would rather be playing SC2.", "and so the crazy refactoring process sees the sunlight after some months in the dark!", "somebody keeps erasing my changes.", "I'm hungry", "FOUTRELIS SUCKS", "TDD: 1, Me: 0", "c&p fail", "Don’t mess with Voodoo", "fuckup.", "Herping the derp derp (silly scoping error)", "That's just how I roll", "Dimitris made me do it", "I don't know what these changes are supposed to accomplish but somebody told me to make them.", "Argh! About to give up :(", "Shovelling coal into the server...", "fixes", "removed tests since i can't make them green", "Well the book was obviously wrong.", "I CAN HAZ PYTHON, I CAN HAZ INDENTS", "more fixes", "This is a basic implementation that works.", "better grepping", "commented out failing tests", "I must have been drunk.", "One little whitespace gets its very own commit! Oh, life is so erratic!", "Ok, 5am, it works.  For real.", "this is how we generate our shit.", "these confounded tests drive me nuts", "just trolling the repo", "done. going to bed now.", "To be honest, I do not quite remember everything I changed here today. But it is all good, I tellya.",
		"Fixed a bug cause Marcus said to", "Blaming regex.", "I'M PUSHING.", "Crap. Tonight is raid night and I am already late.", "Somebody set up us the bomb.", "a few bits tried to escape, but we caught them", "Yep, Rainer was right on this one.", "Reset error count between rows. herpderp", "- Temporary commit.", "final commit.", "Test commit. Please ignore", "wip", "Warun broke the regex, lame", "added some filthy stuff", "This is why git rebase is a horrible horrible thing.", "This is the last time we let Marcus commit ascii porn in the comments.", "Fixed the fuck out of #204!", "small is a real HTML tag, who knew.", "i dunno, maybe this works", "after of this commit remember do a git reset hard", "doh.", "needs more cow bell", "Committing in accordance with the prophecy.", "(O.o)", "(> <) Bunny approves these changes.", "Fixed the fuck out of #322!", "herpderp (redux)", "bug fix", "Too tired to write descriptive message", "I was wrong...", "Derp search replace fuckup", "Fixed errors", "One does not simply merge into master", "Refactored configuration.",
		"fix some fucking errors", "Is there an achievement for this?", "Completed with no bugs...",
		"I know what I am doing. Trust me.", "Fixed a bug cause Tony said to",
		"Code was clean until manager requested to fuck it up", "removed echo and die statements, lolz.", "Issue #9 is now Issue #7", "harharhar", "For great justice.", "various changes", "Fixing Michael's bugs.", "grmbl", "bugger", "Still can't get this right...", "Stephen sucks", "need another beer", "Use a real JS construct, WTF knows why this works in chromium.", "Revert \"just testing, remember to revert\"", "Who knows...", "SEXY RUSSIAN CODES WAITING FOR YOU TO CALL", "unh",
		"Removed test case since code didn't pass QA", "ALL SORTS OF THINGS", "LAST time, foutrelis, /dev/urandom IS NOT a variable name generator...",
		"One more time, but with feeling.", "tl;dr", "pgsql is being a pain", "Todd made me do it", "I was told to leave it alone, but I have this thing called OCD, you see", "It was the best of times, it was the worst of times", "eppic fail Marcus", "Some shit.", "Friday 5pm", "This is the last time we let Chris commit ascii porn in the comments.", "[no message]", "marks", "I'll explain this when I'm sober .. or revert it", "XAVIER, WE WENT OVER THIS. EXPANDTAB.", "Gross hack because Qi doesn't know how to code", "Fixed unnecessary bug.", "grrrr", "HTROYACK, WE WENT OVER THIS. C++ IO SUCKS.", "that coulda been bad", "I must sleep... it's working... in just three hours...", "Nothing to see here, move along", "I don't know what the hell I was thinking.", "lots and lots of changes", "sometimes you just herp the derp so hard it herpderps", "I'm too old for this shit!", "By works, I meant 'doesnt work'.  Works now..", "Arrrrgggg", "It's secret!", "Fixing Andy's bugs.", "FOUTRELIS, WE WENT OVER THIS. EXPANDTAB.", "forgot to save that file", "added super-widget 2.0.", "Same as last commit with changes", "Major fixup.", "Todd rebase plx?", "fail", "Fixed Bug", "Another bug bites the dust", "Who knows WTF?!", "dirty hack, have a better idea ?", "restored deleted entities just to be sure", "Make that it works in 90% of the cases.  3:30.", "it is hump day _^_", "ID:10T Error", "someone fails and it isn't me", "Fixing Robert's bugs.", "Switched off unit test 9 because the build had to go out now and there was no time to fix itproperly.",
		"workaround for ant being a pile of fail", "Don’t even try to refactor it.", "Switched off unit test 6 because the build had to go out now and there was no time to fix itproperly.",
		"Edy rebase plx?", "Gross hack because Pat doesn't know how to code", "You should have trusted me.", "Useful text", "Removed code.", "Gross hack because Alex doesn't know how to code", "Do things better, faster, stronger", "Testing in progress ;)", "I expected something different.", "squash me", "640K ought to be enough for anybody", "should work I guess...", "Switched off unit test 5 because the build had to go out now and there was no time to fix itproperly.",
		"Issue #3 is now Issue #28", "should work now.", "Fixed the fuck out of #543!", "This is why the cat shouldn't sit on my keyboard.", "Fixed the fuck out of #510!", "We should delete this crap before shipping.", "I must enjoy torturing myself", "i think i fixed a bug...", "I hate this fucking language.", "less french words", "This is the last time we let Todd commit ascii porn in the comments.", "It only compiles every 5 tries... good luck.", "eppic fail Jen", "Minor updates", "It only compiles every 3 tries... good luck.", "Jefferson made me do it", "It works!",
	}
)

// Repo general information about repo
type Repo struct {
	DirPath  string
	FilePath string
}

func randomName() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v-%v-%v", adjs[rand.Intn(len(adjs))], nouns[rand.Intn(len(nouns))], repos[rand.Intn(len(repos))])
}

func newRepo(filename string) Repo {
	dirPath := randomName()

	os.Mkdir(dirPath, 0755)
	os.Chdir(dirPath)
	exec.Command("git", "init", ".").Run()

	return Repo{DirPath: dirPath, FilePath: filename}
}

func (r *Repo) appendCommit(data string, date time.Time) {
	err := ioutil.WriteFile(r.FilePath, []byte(data), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Setenv("GIT_AUTHOR_DATE", date.Format(time.RFC3339))
	os.Setenv("GIT_COMMITTER_DATE", date.Format(time.RFC3339))

	exec.Command("git", "add", r.FilePath).Run()
	exec.Command("git", "commit", "-m", messages[rand.Intn(len(messages))]).Run()
}

func main() {
	code := "writeln('Go is Awesome!!!')"
	repo := newRepo("main.go")
	repo.appendCommit(code, time.Now().Add(-24*time.Hour))

	for i := -400; i < 0; i++ {
		d := time.Now().Add(time.Duration(i*24) * time.Hour)
		for j := 0; j < rand.Intn(10); j++ {
			authorDate := time.Date(d.Year(), d.Month(), d.Day(), int(rand.NormFloat64()*3.0+12.0), rand.Intn(59), rand.Intn(59), 0, d.Location())
			uid, err := uuid.NewV5(uuid.NamespaceURL, []byte(time.Now().Format(time.RFC3339Nano)))
			commitData := fmt.Sprintf("%s", uid)
			if err != nil {
				continue
			}
			// fmt.Printf("%v - %v\n", authorDate, commitData)
			repo.appendCommit(commitData, authorDate)
		}
		fmt.Print(".")
	}
	repo.appendCommit(code, time.Now())
	os.Setenv("GIT_AUTHOR_DATE", "")
	os.Setenv("GIT_COMMITTER_DATE", "")

	fmt.Println("\nNow you are a goddamn rockstar!")
}
