<script>
  import { onMount, onDestroy } from "svelte";
  import EditorJS from "@editorjs/editorjs";

  import SimpleImage from "@editorjs/simple-image";
  import Header from "@editorjs/header";
  import List from "@editorjs/list";

  let editor;

  onMount(() => {
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
      autofocus: true,
    });
  });

  // Destroy the editor instance when the component is unmounted
  onDestroy(() => {
    editor?.destroy();
  });
</script>

<div id="editor" />

<style>
  /* Add any desired CSS styles for the editor container */
  #editor {
    border: 1px solid #ccc;
    padding: 10px;
    text-align: left;
  }
</style>
