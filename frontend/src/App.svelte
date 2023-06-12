<script>
  import { onMount } from "svelte";

  import Editor from "./lib/Editor.svelte";
  import { Greet } from "../wailsjs/go/main/App";
  import { main } from "../wailsjs/go/models";

  function generate() {
    let person = new main.Person();
    person.name = "Peter";
    person.age = 27;
    // @ts-ignore
    Greet(person).then((result) => {
      console.log(result);
    });
  }

  generate();
</script>

<main>
  <h1>Note</h1>
  <h3>Take a note, save a note, merge notes</h3>
  <Editor />
</main>
