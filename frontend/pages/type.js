import { useState, useEffect } from "react";

const sentence = "The quick brown fox jumped over the lazy dog";
const tokens = sentence.split(" ");

const Type = () => {
  const [wordCount, setWordCount] = useState(0);
  const [numCorrect, setNumCorrect] = useState(0);
  const [currWord, setCurrWord] = useState("");
  const [correct, setCorrect] = useState(true);

  const keyDown = e => {
    if (e.target.value.slice(-1) === " ") {
      if (correct && currWord === tokens[wordCount]) {
        setNumCorrect(numCorrect + 1);
      }
      // reset
      setCorrect(true);
      setWordCount((wordCount + 1) % tokens.length);
      // TODO: get new words here
      setCurrWord("");
      return;
    }
    setCurrWord(e.target.value);
    const wordCorrect = tokens[wordCount].indexOf(currWord) !== -1;
    setCorrect(wordCorrect);
  };

  return (
    <div>
      <h1>Type!</h1>
      <h3>{currWord}</h3>
      <p>{sentence}</p>
      <p>Number correct: {numCorrect}</p>
      <p>Is correct: {correct}</p>
      <input type="text" value={currWord} onChange={e => keyDown(e)} />
    </div>
  );
};

export default Type;
