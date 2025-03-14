## Deepseek Powered Assitant Guide

### Which Model

Deepseek R1

## Prompting Guides

There is no official guide released by deepseek yet. I found a guide on Deepseek R1 on other website

https://www.learnprompt.org/prompting-deepseek-r1/

## Initial Response

### Deepseek thinking phase:

```
Okay, let's tackle this. The student wants to translate "Did you see the raven this morning? They were looking at our garden." into Japanese. First, I need to break down the sentence into parts.

Start with the vocabulary. The main words are "see" (miru), "raven" (karasu), "this morning" (kesa), "garden" (niwa), and "our" (watashitachi no). I should present these in a table without particles, so the student figures out the particles themselves.

Next, sentence structure. The first question is "Did you see...?" which in Japanese often uses the past tense of miru, which is mimashita. The structure here might be [Time] [Object] [Verb]? For the second sentence, "They were looking at..." which is a continuous action. Since the student is JLPT5, maybe use the te-iru form, so mite imashita. The structure could be [Subject] [Location] [Verb].

Clues and considerations: The student needs to remember to use the past tense for "see" and the past continuous for "looking." Also, "they" might be implied by context, so using karasu-tachi or just leaving it out. The possessive "our" would be watashitachi no niwa. Particles like o for the object in the first sentence and ni for direction in the second.

Possible mistakes: Forgetting to conjugate miru to mimashita, incorrect particle usage (o vs. ga), and forming the continuous tense correctly. Also, ordering of time phrases—kesa should come at the beginning.

When the student attempts, check if they used the correct verb forms and particles. If they make an error, give hints about tense or particles without giving away the answer. Encourage them to refer to the vocabulary table and structure examples.
```

Vocabulary Table

| Japanese	| Romaji    | English |
|-----------|-----------|---------|
|見る       | miru	    |   to see|
|カラス	    | karasu	|raven|
|今朝	    |kesa	    |this morning|
|庭	        | niwa      |garden |
|私たちの	 |watashitachi |no	our|

Sentence Structure

First Sentence: [Time] [Object] [Verb (past tense)]?
Example: Did you see the raven this morning? → [今朝] [カラス] [見る]?

Second Sentence: [Subject (implied "they")] [Location] [Verb (continuous past tense)].
Example: They were looking at our garden. → [私たちの庭] [見る].

Clues & Considerations

Use the past tense for "see" (見ました mimashita).

"They" can be omitted in Japanese if contextually clear.

"Looking at" requires the continuous form (見ていました mite imashita).

"Our garden" uses the possessive 私たちの庭.

Next Steps:

Try constructing the first sentence using the structure above.

Remember particles like を (object marker) and に (direction marker).

<hr>

## Response after providing xml files 

### Deepseek thinking phase:

```
Okay, let's break down how to approach translating the sentence "Did you see the raven this morning? They were looking at our garden." into Japanese based on the provided examples and guidelines.

First, I'll start by analyzing the sentence structure. The first part is a question: "Did you see the raven this morning?" According to the sentence-structure-examples.xml, a similar structure is "[Time] [Subject] [Object] [Verb]?" For example, "This morning, I saw the butterfly." becomes "[Time] [Subject] [Object] [Verb]." But since this is a question, the structure might be adjusted to "[Time] [Object] [Verb]?" with the subject implied as "you."

Next, the vocabulary needed includes "raven," "see," "this morning," "they," "look at," and "garden." Checking the examples in example.xml, the vocabulary table should list the Japanese terms with their romaji and English equivalents. It's important to use dictionary forms and avoid conjugations here.

For the verb "see," the dictionary form is 見る (miru). "Raven" is 烏 (karasu). "This morning" is 今朝 (kesa). "Garden" is 庭 (niwa). "They" can be translated as 彼ら (karera), but in Japanese, pronouns are often omitted if contextually clear. "Looking at" would use the verb 見る with the particle を (wo) for the object and を (wo) or に (ni) for the direction, but since the example in considerations-examples.xml uses を見る (wo miru) for "looking at," that might be appropriate here.

The second sentence, "They were looking at our garden," is a statement. The structure here would be "[Subject] [Object] [Verb]." Since "they" can be omitted, the focus is on the garden and the action of looking. The past progressive tense "were looking" would use the ～ていた (te ita) form. So 見ていた (mite ita).

Putting it together, the first question would be structured as "今朝、烏を見ましたか？" (Kesa, karasu wo mimashita ka?). The second sentence would be "庭を見ていました。" (Niwa wo mite imashita.) However, considering politeness levels, since the user might need to use plain form based on the clues not giving away polite forms, it might be 見ていた (mite ita) instead of 見ていました (mite imashita).

Checking the scoring criteria from example.xml, the vocabulary table should be concise, structure clear, and clues not giving away conjugations. The assistant's output should start with the vocabulary table, then the structure, considerations, and next steps.

Potential considerations include:

- Using を particle for the direct object (raven and garden).

- Time placement at the beginning of the sentence.

- Handling past tense and progressive aspect in Japanese.

- Omitting pronouns when contextually clear.

Possible next steps for the student could be attempting the sentence, asking about particle usage, or verb conjugation for past progressive.
```

Cheatsheet:  

| Word          | Japanese       | Type                   |  
|---------------|----------------|------------------------|  
| raven         | 烏（からす）    | Noun                   |  
| see           | 見る（みる）    | Transitive Ichidan Verb|  
| this morning  | 今朝（けさ）    | Noun (Time)            |  
| garden        | 庭（にわ）      | Noun                   |  
| to look at    | 見る（みる）    | Transitive Ichidan Verb|  

Sentence structure:  
1. **[Time] [Object] [Verb-past]?**  
   *(Question about past action with time context)*  
2. **[Object] [Verb-past-progressive].**  
   *(Statement about ongoing past action, subject omitted)*  

Considerations:  
- Use を (wo) particle to mark objects ("raven" and "garden").  
- Time phrase (今朝) typically comes first.  
- Past tense for "see" → 見ました (mimashita) or 見た (mita).  
- Past progressive "were looking" → 見ていた (mite ita).  
- Subject omission in Japanese ("you" and "they" may be dropped).  

Next steps:  
- Attempt to write the full sentence.  
- Ask for help with:  
  - Particle usage (を vs. に).  
  - Verb conjugation for past progressive.  
  - How to connect two sentences naturally.