<script>
  import Editor from "./lib/Editor.svelte";
  import { Greet } from "../wailsjs/go/main/App";
  import { main } from "../wailsjs/go/models";

  let person = new main.Person();
  person.name = "Peter";
  person.age = 27;

  const handleInput = (event) => {
    const val = event.target.value;
    person.name = val;

    // @ts-ignore
    Greet(person).then((result) => {
      console.log(result);
    });
  };
</script>

<main>
  <h1>Note</h1>
  <h3>Take a note, save a note, merge notes</h3>
  <Editor />
  <input placeholder="What are your thoughts today?" on:input={handleInput} />

  <h1>Hello {person.name}!</h1>
</main>
