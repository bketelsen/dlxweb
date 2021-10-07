<script>
  // can be one of light or dark
  export let color = "light";

  import { onMount } from "svelte";
  import { InstanceService } from '../../oto';

  let containers = [];

  onMount(async () => {
    const instanceService = new InstanceService();
    const result = await instanceService.list();
    containers = result.instances

  });
  function getClass(status) {
    var style = "fas fa-circle "
    if (status === "Running") {
      style = style + "text-green-500"
    } else if (status === "Stopped") {
      style = style + "text-red-500"
    }
    style = style + " mr-2"
    return style
  }
</script>

<div
  class="relative flex flex-col min-w-0 break-words w-full mb-6 shadow-lg rounded {color ===
  'light'
    ? 'bg-white'
    : 'bg-red-800 text-white'}"
>
  <div class="rounded-t mb-0 px-4 py-3 border-0">
    <div class="flex flex-wrap items-center">
      <div class="relative w-full px-4 max-w-full flex-grow flex-1">
        <h3
          class="font-semibold text-lg {color === 'light'
            ? 'text-gray-800'
            : 'text-white'}"
        >
          Containers
        </h3>
      </div>
    </div>
  </div>
  <div class="block w-full overflow-x-auto">
    <!-- Projects table -->
    <table class="items-center w-full bg-transparent border-collapse">
      <thead>
        <tr>
          <th
            class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-no-wrap font-semibold text-left {color ===
            'light'
              ? 'bg-gray-100 text-gray-600 border-gray-200'
              : 'bg-red-700 text-red-200 border-red-600'}"
          >
            Name
          </th>
          <th
            class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-no-wrap font-semibold text-left {color ===
            'light'
              ? 'bg-gray-100 text-gray-600 border-gray-200'
              : 'bg-red-700 text-red-200 border-red-600'}"
          >
            IP Address
          </th>
          <th
            class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-no-wrap font-semibold text-left {color ===
            'light'
              ? 'bg-gray-100 text-gray-600 border-gray-200'
              : 'bg-red-700 text-red-200 border-red-600'}"
          >
            Status
          </th>
          
        </tr>
      </thead>
      <tbody>
        {#each containers as container}
        <tr>
          <td
            class="border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-no-wrap p-4 text-left flex items-center"
          >
            <span
              class="font-bold {color === 'light'
                ? 'text-gray-700'
                : 'text-white'}"
            >
              {container.name}
            </span>
          </td>
          <td
            class="border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-no-wrap p-4"
          >
            {container.iPV4}
          </td>
          <td
            class="border-t-0 px-6 align-middle border-l-0 border-r-0 text-xs whitespace-no-wrap p-4"
          >
            <i class={getClass(container.status)}/> {container.status}
          </td>

        </tr>
        {:else}
        <!-- this block renders when photos.length === 0 -->
        <p>loading...</p>
        {/each}

      </tbody>
    </table>
  </div>
</div>
