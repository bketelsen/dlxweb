<script>
  // library for creating dropdown menu appear on click
  import { createPopper } from "@popperjs/core";
  import { project } from '../../stores/project.js'

  // core components
  export let projects = [];
  let dropdownPopoverShow = false;

  let btnDropdownRef;
  let popoverDropdownRef;

  const toggleDropdown = () => {
    if (dropdownPopoverShow) {
      dropdownPopoverShow = false;
    } else {
      dropdownPopoverShow = true;
      createPopper(btnDropdownRef, popoverDropdownRef, {
        placement: "bottom-start",
      });
    }
  };

  function setProject(e) {
    console.log("setting", e);
	project.set(e.target.innerText.trim())
	toggleDropdown()
  }
</script>

<div class="flex flex-wrap">
  <div class="w-full px-4">
    <div class="relative inline-flex align-middle w-full">
      <button
        class="text-white font-bold uppercase text-sm px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 bg-gray-800 active:bg-gray-900 ease-linear transition-all duration-150"
        type="button"
        bind:this={btnDropdownRef}
        on:click={toggleDropdown}
      >
        Project ({$project})
      </button>
      <div
        bind:this={popoverDropdownRef}
        class="bg-white  text-base z-50 float-left py-2 list-none text-left rounded shadow-lg mt-1 min-w-48 {dropdownPopoverShow
          ? 'block'
          : 'hidden'}"
      >
        {#each projects as project}
          <span
            on:click={setProject}
            class="text-sm py-2 px-4 font-normal block w-full whitespace-no-wrap bg-transparent text-gray-800"
          >
            {project.name}
          </span>
        {:else}
          <span>Loading</span>
        {/each}
      </div>
    </div>
  </div>
</div>
