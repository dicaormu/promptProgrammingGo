package practices

import "strings"


type Practices struct{
   Text string
   Prompt string
}

//Use delimiters to clearly indicate distinct parts of the input
//Delimiters can be anything like: ```, """, < >, <tag> </tag>, :

func OnlyDelimiters() Practices{
 return Practices{
	Text: "You should express what you want a model to do by "+
	"providing instructions that are as clear and " +
	"specific as you can possibly make them. " +
	"This will guide the model towards the desired output, " +
	"and reduce the chances of receiving irrelevant " +
	"or incorrect responses. Don't confuse writing a " +
	"clear prompt with writing a short prompt. " +
	"In many cases, longer prompts provide more clarity " +
	"and context for the model, which can lead to " +
	"more detailed and relevant outputs.",
	Prompt: "Summarize the text delimited by triple backticks into a single sentence. ```{Text}```" }
}

//Tactic 2: Ask for a structured output
func StructuredOutput() Practices {
	return Practices{
		Text: "Generate a list of ingredients required to prepare a medium regina pizza for 4 people "+
		"with amounts and an estimated price. Provide them in JSON format with the following keys: "+
		"ingredient, amount, price",
		Prompt: "```{Text}```"}
}

//Tactic 3: Ask the model to check whether conditions are satisfied
func AreConditionsSatisfied() Practices {
	return Practices{
		Text: "Preheat the oven to 240°C. Cover a plate with baking paper. Put 250 ml of lukewarm water and the "+
		"yeast in the bowl fitted with kneading/grinding tool. Start the pastry program P1. After 30 secs, add the flour, "+
		"salt and 2 tbsp of olive oil. "+
		"Cut the slices of ham into four. Chop the mushrooms into thin slices and cut the mozzarella into cubes. "+
		"At the end of the program, roll out the dough on a plate covered with baking paper. "+
		"Spread the tomato purée on the dough and sprinkle with oregano. Add the mozzarella, ham and mushrooms."+
		"Sprinkle with a little oil. Bake for approximately 15 mins. "+
		"TIP Choose the topping according to your preference",
		Prompt: "You will be provided with text delimited by triple backticks. "+
		"If it contains a sequence of instructions, "+ 
		"re-write those instructions in the following format: \\"+
		
		"Step 1 - ... \\"+
		"Step 2 - ...\\"+
		"...\\"+
		"Step N - ...\\"+
		"If the text does not contain a sequence of instructions, "+
		"then simply write \"No steps provided.\" "+
		"```{Text}```" }
}


// Tactic 4: "Few-shot" prompting
func FewShotPrompting() Practices {
	return Practices{
		Text: "",
		Prompt: "Your task is to answer in a consistent style. "+
		"<child>: Teach me about patience."+
		"<grandparent>: The river that carves the deepest "+
		"valley flows from a modest spring; the "+
		"grandest symphony originates from a single note; "+
		"the most intricate tapestry begins with a solitary thread."+
		"<child>: Teach me about resilience."}
}

//Principle 2: Give the model time to “think”
//Tactic 1: Specify the steps required to complete a task
func SpecifySteps() Practices {
	return Practices{
		Text: "Hansel and Gretel are the young children of a poor woodcutter. When a famine settles over the land, the woodcutter's second wife tells the woodcutter to take the children into the woods and leave them there to fend for themselves, so that she and her husband do not starve to death as the children, particularly Hansel, eat far too much food. The woodcutter opposes the plan, but his wife claims that maybe a stranger will take the children in and provide for them, which the woodcutter and she simply cannot do. With the scheme seemingly justified, the woodcutter reluctantly is forced to submit to it. They are unaware that in the children's bedroom, Hansel and Gretel have overheard them. After the parents have gone to bed, Hansel sneaks out of the house and gathers as many white pebbles as he can, then returns to his room, reassuring Gretel that God will not forsake them",
		Prompt:"Perform the following actions: "+
		"1 - Summarize the following text delimited by triple backticks with 1 sentence. "+
		"2 - Translate the summary into French. "+
		"3 - List each name in the French summary. "+
		"4 - Output a json object that contains the following "+
		"keys: french_summary, num_names. "+
		"Separate your answers with line breaks. "+
		"Text: "+
		"```{Text}```"}
}

// Tactic 2: Instruct the model to work out its own solution before rushing to a conclusion
func WorkOwnSolution() Practices {
	return Practices{
		Text: "",
		Prompt: "Determine if the student's solution is correct or not. "+
		"Your task is to determine if the student's solution "+
		"is correct or not. "+
		"To solve the problem do the following: "+
		"- First, work out your own solution to the problem. "+
		"- Then compare your solution to the student's solution "+
		"and evaluate if the student's solution is correct or not. "+
		"Don't decide if the student's solution is correct until "+
		"you have done the problem yourself. "+
		"Use the following format: "+
		"Question: "+
		"``` "+
		"question here "+
		"``` "+
		"Student solution: "+
		"``` "+
		"student solution here "+
		"``` " +
		"Actual solution: "+
		"``` "+
		"steps to work out the solution and your solution here "+
		"``` "+
		"Is the student solution the same as actual solution "+
		"just calculated: "+
		"``` "+
		"yes or no "+
		"``` "+
		"Student grade: "+
		"``` "+
		"correct or incorrect "+
		"``` "+
	"Question: "+
	"I'm building a solar power installation and I need "+
	" help working out the financials. "+
	"- Land costs $100 / square foot "+
	"- I can buy solar panels for $250 / square foot "+
	"- I negotiated a contract for maintenance that will cost "+
	"me a flat $100k per year, and an additional $10 / square "+
	"foot "+
	"What is the total cost for the first year of operations "+
	"as a function of the number of square feet. "+
	"Student Solution: "+
	"Let x be the size of the installation in square feet. "+
	"Costs: "+ 
	"1. Land cost: 100x "+ 
	"2. Solar panel cost: 250x "+
	"3. Maintenance cost: 100,000 + 100x "+
	"Total cost: 100x + 250x + 100,000 + 100x = 450x + 100,000" }
}

// limitations: hallucinations
func Hallucination() Practices {
	return Practices{
		Text: "",
		Prompt: "Tell me about hungerstation api for delivery with chatgpt "}
	}

func ParsePractice(practice Practices) string{
		return strings.Replace(practice.Prompt,"{Text}",practice.Text,1)
}