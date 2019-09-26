import { useState, useEffect } from "react";
import { WordProvider } from "my-utils";

// const sentence = "The quick brown fox jumped over the lazy dog";
// const tokens = sentence.split(" ");

const wp = new WordProvider();
const Type = () => {
  const [provider, _] = useState(wp);
  const [tokens, setTokens] = useState([]);
  const [wordCount, setWordCount] = useState(0);
  const [numCorrect, setNumCorrect] = useState(0);
  const [currWord, setCurrWord] = useState("");
  const [correct, setCorrect] = useState(true);
  const [shouldGetMoreWords, setShouldGetMoreWords] = useState(true);

  useEffect(() => {
    async function getMoreWords() {
      const toks = await provider.getMoreWords.bind(provider)();
      setTokens(toks);
      setShouldGetMoreWords(false);
    }
    if (shouldGetMoreWords) {
      getMoreWords();
    }
  }, [shouldGetMoreWords]);

  const keyDown = e => {
    if (e.target.value.slice(-1) === " ") {
      if (correct && currWord === tokens[wordCount]) {
        setNumCorrect(numCorrect + 1);
      }
      // reset
      setCorrect(true);

      setWordCount((wordCount + 1) % tokens.length);
      // useEffect will fetch new words on next render
      setShouldGetMoreWords(true);
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
      <ul>
        {tokens.map(t => (
          <li>{t}</li>
        ))}
      </ul>
      <p>Number correct: {numCorrect}</p>
      <p>Is correct: {correct}</p>
      <input type="text" value={currWord} onChange={e => keyDown(e)} />
    </div>
  );
};

export default Type;
