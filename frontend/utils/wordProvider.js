const fetch = require("isomorphic-unfetch");
const BACKEND_HOST = "http://localhost:5000";

/**
 * WordProvider is a class that is responsible for fetching words for competitors to type
 */
class WordProvider {
  constructor() {
    this.words = [];
  }

  async getMoreWords() {
    const res = await fetch(`${BACKEND_HOST}/words`);
    if (res.status !== 200) {
      console.error("Request to get more words failed!");
      // winston here
    } else {
      console.log({ this: this });
      this.words.push(res.data.words);
    }
    return this.words;
  }
}

module.exports = WordProvider;
