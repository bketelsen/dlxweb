<script>
  // can be one of light or dark
  export let color = "light";
  import { toast } from "@zerodevx/svelte-toast";

  import { InstanceService } from "../../oto";
  import { project } from "../../stores/project";

  import ContainerControls from "components/Controls/ContainerControls.svelte";
  let containers = [];

  let selectedProject;
  let newContainerName;

  // Create Container modal
  let showModal = false;

  function toggleModal() {
    showModal = !showModal;
  }

  const instanceService = new InstanceService();
  const unsubscribe = project.subscribe((value) => {
    selectedProject = value;
  });

  async function getInstances(proj) {
    const result = await instanceService.list({ project: proj });
    containers = result.instances;
    return containers;
  }

  $: containerPromise = getInstances(selectedProject);

  function getClass(status) {
    var style = "fas fa-circle ";
    if (status === "Running") {
      style = style + "text-green-500";
    } else if (status === "Stopped") {
      style = style + "text-red-500";
    }
    style = style + " mr-2";
    return style;
  }

  function refreshInstances() {
    containerPromise = getInstances(selectedProject);
  }

  function handleCreate(event) {
    if (newContainerName) {
      const container = {
        name: newContainerName,
        project: selectedProject,
      };
      const id = toast.push({
        msg: "<i class='fas fa-info'></i> Creating " + newContainerName,
      });
      instanceService.create(container).then(() => {
        const newid = toast.push({
          msg: "<i class='fas fa-info'></i> Created " + newContainerName,
        });
        toggleModal();
        toast.pop(id)
        containerPromise = getInstances(selectedProject);

      });
    } else {
      const id = toast.push({
        msg: "<i class='fas fa-info'></i> Container Name Required",
      });
    }
  }

  function handleStart(event) {
    const id = toast.push({
      msg: "<i class='fas fa-info'></i> Starting " + event.detail.container,
    });
    instanceService
      .start({
        project: selectedProject,
        name: event.detail.container,
      })
      .then(() => {
        console.log("Started");
        containerPromise = getInstances(selectedProject);
      });
  }
  function handleStop(event) {
    const id = toast.push({
      msg: "<i class='fas fa-info'></i> Stopping " + event.detail.container,
    });
    instanceService
      .stop({
        project: selectedProject,
        name: event.detail.container,
      })
      .then(() => {
        console.log("Stopped");
        containerPromise = getInstances(selectedProject);
      });
  }
    function handleDelete(event) {
    const id = toast.push({
      msg: "<i class='fas fa-info'></i> Removing " + event.detail.container,
    });
    instanceService
      .delete({
        project: selectedProject,
        name: event.detail.container,
      })
      .then(() => {
        console.log("deleted");
        containerPromise = getInstances(selectedProject);
      });
      
  }
</script>

<div
  class="relative flex flex-col min-w-0 break-words w-full mb-6 shadow-lg rounded {color ===
  'light'
    ? 'bg-white'
    : 'bg-red-800 text-white'}"
>
  <div class="px-4 py-3 mb-0 border-0 rounded-t">
    <div class="flex flex-wrap items-center">
      <div class="relative flex-1 flex-grow w-full max-w-full px-4">
        <h3
          class="font-semibold text-lg {color === 'light'
            ? 'text-gray-800'
            : 'text-white'}"
        >
          Containers
        </h3>
      </div>
      <div class="relative flex-1 flex-grow w-full max-w-full px-4">
        <button
          class="px-4 py-2 mb-1 mr-1 text-xs font-bold text-white uppercase transition-all duration-150 ease-linear bg-blue-500 rounded shadow outline-none active:bg-blue-700 hover:shadow-md focus:outline-none"
          type="button"
          on:click={() => toggleModal()}
        >
          New Container
        </button>
        <button
        class="px-4 py-2 mb-1 mr-1 text-xs font-bold text-white uppercase transition-all duration-150 ease-linear bg-blue-500 rounded shadow outline-none active:bg-blue-700 hover:shadow-md focus:outline-none"
        type="button"
        on:click={() => refreshInstances()}
      >
        <i class="fas fa-sync-alt"></i>
      </button>
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

          <th
            class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-no-wrap font-semibold text-left {color ===
            'light'
              ? 'bg-gray-100 text-gray-600 border-gray-200'
              : 'bg-red-700 text-red-200 border-red-600'}"
          >
            Action
          </th>
        </tr>
      </thead>
      <tbody>
        {#await containerPromise then containers}
          {#each containers as container}
            <tr>
              <td
                class="flex items-center p-4 px-6 text-xs text-left whitespace-no-wrap align-middle border-t-0 border-l-0 border-r-0"
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
                class="p-4 px-6 text-xs whitespace-no-wrap align-middle border-t-0 border-l-0 border-r-0"
              >
                {container.iPV4}
              </td>
              <td
                class="p-4 px-6 text-xs whitespace-no-wrap align-middle border-t-0 border-l-0 border-r-0"
              >
                <i class={getClass(container.status)} />
                {container.status}
              </td>
              <td
                class="p-4 px-6 text-xs whitespace-no-wrap align-middle border-t-0 border-l-0 border-r-0"
              >
                <ContainerControls
                  {container}
                  on:start={handleStart}
                  on:stop={handleStop}
                  on:delete={handleDelete}
                />
              </td>
            </tr>
          {:else}
            <!-- this block renders when photos.length === 0 -->
            <p>loading...</p>
          {/each}
        {/await}
      </tbody>
    </table>
  </div>
</div>
{#if showModal}
  <div
    class="fixed inset-0 z-50 flex items-center justify-center overflow-x-hidden overflow-y-auto outline-none focus:outline-none"
  >
    <div class="relative w-auto max-w-sm mx-auto my-6">
      <!--content-->
      <div
        class="relative flex flex-col w-full bg-white border-0 rounded-lg shadow-lg outline-none focus:outline-none"
      >
        <!--header-->
        <div
          class="flex items-start justify-between p-5 border-b border-solid rounded-t border-blueGray-200"
        >
          <h3 class="text-3xl font-semibold">New Container</h3>
          <button
            class="float-right p-1 ml-auto text-3xl font-semibold leading-none text-black bg-transparent border-0 outline-none opacity-5 focus:outline-none"
            on:click={toggleModal}
          >
            <span
              class="block w-6 h-6 text-2xl text-black bg-transparent outline-none opacity-5 focus:outline-none"
            >
              ×
            </span>
          </button>
        </div>
        <!--body-->
        <div class="relative flex-auto p-6">
          <div class="pt-0 mb-3">
            <h4 class="text-2xl font-semibold">Container Name</h4>
            <input
              bind:value={newContainerName}
              type="text"
              placeholder="name"
              class="relative w-full px-3 py-3 text-sm bg-white rounded shadow outline-none placeholder-blueGray-300 text-blueGray-600 focus:outline-none focus:shadow-outline"
            />
          </div>
        </div>
        <!--footer-->
        <div
          class="flex items-center justify-end p-6 border-t border-solid rounded-b border-blueGray-200"
        >
          <button
            class="px-6 py-2 mb-1 mr-1 text-sm font-bold text-red-500 uppercase transition-all duration-150 ease-linear outline-none background-transparent focus:outline-none"
            type="button"
            on:click={toggleModal}
          >
            Close
          </button>
          <button
            class="px-6 py-3 mb-1 mr-1 text-sm font-bold text-black uppercase transition-all duration-150 ease-linear rounded shadow outline-none bg-emerald-500 active:bg-emerald-600 hover:shadow-lg focus:outline-none"
            type="button"
            on:click={handleCreate}
          >
            Create Container
          </button>
        </div>
      </div>
    </div>
  </div>
  <div class="fixed inset-0 z-40 bg-black opacity-25" />
{/if}
