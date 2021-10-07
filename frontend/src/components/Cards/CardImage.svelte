<script>
  import { ImageService } from "../../oto";
  import { onMount } from "svelte";

  let source = ""
  let username = "bjk"

  onMount(async () => {
    const imageService = new ImageService();
    const result = await imageService.source();
    source = result.source

  });

  function changeUsername(event) {
    
    const newsource = source.replaceAll('dlxuser', username)
    source = newsource
  }
</script>

<div
  class="relative flex flex-col w-full min-w-0 mb-6 break-words bg-gray-200 border-0 rounded-lg shadow-lg"
>
  <div class="px-6 py-6 mb-0 bg-white rounded-t">
    <div class="flex justify-between text-center">
      <h6 class="text-xl font-bold text-gray-800">Base Image</h6>
      <button
        class="px-4 py-2 mr-1 text-xs font-bold text-white uppercase transition-all duration-150 ease-linear bg-red-500 rounded shadow outline-none active:bg-red-600 hover:shadow-md focus:outline-none"
        type="button"
      >
        GENERATE IMAGE
      </button>
    </div>
  </div>
  <div class="flex-auto px-4 py-10 pt-0 lg:px-10">
    <form>
      <h6 class="mt-3 mb-6 text-sm font-bold text-gray-500 uppercase">
        User Information
      </h6>
      <div class="flex flex-wrap">
        <div class="w-full px-4 lg:w-6/12">
          <div class="relative w-full mb-3">
            <label
              class="block mb-2 text-xs font-bold text-gray-700 uppercase"
              for="grid-username"
            >
              Linux Username
            </label>
            <input
              id="grid-username"
              type="text"
              class="px-3 py-3 text-sm text-gray-700 placeholder-gray-400 transition-all duration-150 ease-linear bg-white rounded shadow focus:outline-none focus:shadow-outline"
              bind:value={username}
              on:click={() => changeUsername()}
            />
            <button
            class="px-4 py-2 mr-1 text-xs font-bold text-white uppercase transition-all duration-150 ease-linear bg-red-500 rounded shadow outline-none active:bg-red-600 hover:shadow-md focus:outline-none"
            type="button"
          >
            Apply
          </button>
          </div>
        </div>
      </div>


      <hr class="mt-6 border-gray-400 border-b-1" />

      <h6 class="mt-3 mb-6 text-sm font-bold text-gray-500 uppercase">
        Source
      </h6>
      <div class="flex flex-wrap">
        <div class="w-full px-4 lg:w-12/12">
          <div class="relative w-full mb-3">
            <label
              class="block mb-2 text-xs font-bold text-gray-700 uppercase"
              for="grid-about-me"
            >
              dlxbase.yaml
            </label>
            <textarea
              id="grid-about-me"
              type="text"
              class="w-full px-3 py-3 text-sm text-gray-700 placeholder-gray-400 transition-all duration-150 ease-linear bg-white rounded shadow focus:outline-none focus:shadow-outline"
              rows="40"
              bind:value={source}
            />
          </div>
        </div>
      </div>
    </form>
  </div>
</div>
