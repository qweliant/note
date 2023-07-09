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
      data: {
        time: 1550476186479,
        blocks: [
          {
            id: "oUq2g_tl8y",
            type: "header",
            data: {
              text: "Editor.js",
              level: 2,
            },
          },
          {
            id: "zbGZFPM-iI",
            type: "paragraph",
            data: {
              text: "Hey. Meet the new Editor. On this page you can see it in action — try to edit this text. Source code of the page contains the example of connection and configuration.",
            },
          },
          {
            id: "qYIGsjS5rt",
            type: "header",
            data: {
              text: "Key features",
              level: 3,
            },
          },
          {
            id: "XV87kJS_H1",
            type: "list",
            data: {
              style: "unordered",
              items: [
                "It is a block-styled editor",
                "It returns clean data output in JSON",
                "Designed to be extendable and pluggable with a simple API",
              ],
            },
          },
          {
            id: "AOulAjL8XM",
            type: "header",
            data: {
              text: "What does it mean «block-styled editor»",
              level: 3,
            },
          },
          {
            id: "cyZjplMOZ0",
            type: "paragraph",
            data: {
              text: 'Workspace in classic editors is made of a single contenteditable element, used to create different HTML markups. Editor.js <mark class="cdx-marker">workspace consists of separate Blocks: paragraphs, headings, images, lists, quotes, etc</mark>. Each of them is an independent contenteditable element (or more complex structure) provided by Plugin and united by Editor\'s Core.',
            },
          },
        ],
        version: "2.8.1",
      },
    });
  });

  const save = async () => {
    const outputData = await editor.save();
    console.log("Article data: ", outputData);
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
