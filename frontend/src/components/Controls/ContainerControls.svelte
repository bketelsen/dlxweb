<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  export let container;

  function getStartClass(status) {
    var style = "fas fa-play ";
    if (status === "Running") {
      style = style + "text-gray-500";
    } else if (status === "Stopped") {
      style = style + "text-green-500";
    }
    style = style + " mr-2";
    return style;
  }

  function getStopClass(status) {
    var style = "fas fa-square ";
    if (status === "Running") {
      style = style + "text-red-500";
    } else if (status === "Stopped") {
      style = style + "text-gray-500";
    }
    style = style + " mr-2";
    return style;
  }

    function getDeleteClass(status) {
    var style = "fas fa-trash ";
    if (status === "Running") {
      style = style + "text-gray-500";
    } else if (status === "Stopped") {
      style = style + "text-black-500";
    }
    style = style + " mr-2";
    return style;
  }

  function stopContainer() {
    dispatch("stop", {
      container: container.name,
    });
  }

  function startContainer() {
    dispatch("start", {
      container: container.name,
    });
  }
  function deleteContainer() {
    if (container.status === "Running") {
      alert("Container is running, please stop it first");
      return;
    }
    dispatch("delete", {
      container: container.name,
    });
  }
</script>

<div>
  <i class={getStartClass(container.status)} on:click={startContainer} />
  <i class={getStopClass(container.status)} on:click={stopContainer} />
  <i class={getDeleteClass(container.status)} on:click={deleteContainer} />
</div>
