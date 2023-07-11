<script lang="ts">
  import { onMount, onDestroy } from "svelte";
  import EditorJS, { type OutputData } from "@editorjs/editorjs";

  import {
    SaveNote,
    WriteProtoToFile,
    ReadProtoFromFile,
  } from "../../wailsjs/go/app/App";
  import { app } from "../../wailsjs/go/models";

  import SimpleImage from "@editorjs/simple-image";
  import Header from "@editorjs/header";
  import List from "@editorjs/list";

  let note = new app.Note();
  let editor: EditorJS;

  onMount(async () => {
    const note = await ReadProtoFromFile("data/note.bin");
    const initialData = note.body ? (JSON.parse(note.body) as OutputData) : undefined;
    // Initialize the EditorJS instance
    editor = new EditorJS({
      holder: "editor",

      // Configure the desired tools
      tools: {
        header: {
          class: Header,
          inlineToolbar: true,
        },
        list: {
          class: List,
          inlineToolbar: true,
        },
        image: SimpleImage,
      },
      data: initialData,
    });
  });

  const save = async () => {
    const outputData = await editor.save();
    note.body = JSON.stringify(outputData);
    SaveNote(note).then((result) => {
      console.log(result);
    });
    WriteProtoToFile("data/note.bin", note).then((result) => {
      console.log(result);
    });
  };

  // Destroy the editor instance when the component is unmounted
  onDestroy(() => {
    editor?.destroy();
  });
</script>

<div id="editor">
  <button on:click={save}>Save</button>
</div>

<style>
  /* Add any desired CSS styles for the editor container */
  #editor {
    border: 1px solid #ccc;
    padding: 10px;
    text-align: left;
  }
</style>
