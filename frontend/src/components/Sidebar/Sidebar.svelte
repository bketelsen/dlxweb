<script>
  import { link } from "svelte-routing";

  // core components
  import NotificationDropdown from "components/Dropdowns/NotificationDropdown.svelte";
  import UserDropdown from "components/Dropdowns/UserDropdown.svelte";

  let collapseShow = "hidden";

  function toggleCollapseShow(classes) {
    collapseShow = classes;
  }

  export let location;
</script>

<nav
  class="md:left-0 md:block md:fixed md:top-0 md:bottom-0 md:overflow-y-auto md:flex-row md:flex-no-wrap md:overflow-hidden shadow-xl bg-white flex flex-wrap items-center justify-between relative md:w-64 z-10 py-4 px-6"
>
  <div
    class="md:flex-col md:items-stretch md:min-h-full md:flex-no-wrap px-0 flex flex-wrap items-center justify-between w-full mx-auto"
  >
    <!-- Toggler -->
    <button
      class="cursor-pointer text-black opacity-50 md:hidden px-3 py-1 text-xl leading-none bg-transparent rounded border border-solid border-transparent"
      type="button"
      on:click={() => toggleCollapseShow("bg-white m-2 py-3 px-6")}
    >
      <i class="fas fa-bars" />
    </button>
    <!-- Brand -->
    <a
      use:link
      class="md:block text-left md:pb-2 text-gray-700 mr-0 inline-block whitespace-no-wrap text-sm uppercase font-bold p-4 px-0"
      href="/"
    >
      DLX
    </a>
    <!-- User -->
    <ul class="md:hidden items-center flex flex-wrap list-none">
      <li class="inline-block relative">
        <NotificationDropdown />
      </li>
      <li class="inline-block relative">
        <UserDropdown />
      </li>
    </ul>
    <!-- Collapse -->
    <div
      class="md:flex md:flex-col md:items-stretch md:opacity-100 md:relative md:mt-4 md:shadow-none shadow absolute top-0 left-0 right-0 z-40 overflow-y-auto overflow-x-hidden h-auto items-center flex-1 rounded {collapseShow}"
    >
      <!-- Collapse header -->
      <div
        class="md:min-w-full md:hidden block pb-4 mb-4 border-b border-solid border-gray-300"
      >
        <div class="flex flex-wrap">
          <div class="w-6/12">
            <a
              use:link
              class="md:block text-left md:pb-2 text-gray-700 mr-0 inline-block whitespace-no-wrap text-sm uppercase font-bold p-4 px-0"
              href="/"
            >
              DLX
            </a>
          </div>
          <div class="w-6/12 flex justify-end">
            <button
              type="button"
              class="cursor-pointer text-black opacity-50 md:hidden px-3 py-1 text-xl leading-none bg-transparent rounded border border-solid border-transparent"
              on:click={() => toggleCollapseShow("hidden")}
            >
              <i class="fas fa-times" />
            </button>
          </div>
        </div>
      </div>
      <!-- Form -->
      <form class="mt-6 mb-4 md:hidden">
        <div class="mb-3 pt-0">
          <input
            type="text"
            placeholder="Search"
            class="px-3 py-2 h-12 border border-solid border-gray-600 placeholder-gray-400 text-gray-700 bg-white rounded text-base leading-snug shadow-none outline-none focus:outline-none w-full font-normal"
          />
        </div>
      </form>

      <!-- Divider -->
      <hr class="my-4 md:min-w-full" />
      <!-- Heading -->
      <h6
        class="md:min-w-full text-gray-600 text-xs uppercase font-bold block pt-1 pb-4 no-underline"
      >
        LXD Administration
      </h6>
      <!-- Navigation -->

      <ul class="md:flex-col md:min-w-full flex flex-col list-none">
        <li class="items-center">
          <a
            use:link
            href="/dashboard"
            class="text-xs uppercase py-3 font-bold block {location.href.indexOf(
              '/dashboard'
            ) !== -1
              ? 'text-red-500 hover:text-red-600'
              : 'text-gray-800 hover:text-gray-600'}"
          >
            <i
              class="fas fa-tv mr-2 text-sm {location.href.indexOf(
                '/dashboard'
              ) !== -1
                ? 'opacity-75'
                : 'text-gray-400'}"
            />
            Dashboard
          </a>
        </li>

        <li class="items-center">
          <a
            use:link
            href="/containers"
            class="text-xs uppercase py-3 font-bold block {location.href.indexOf(
              '/containers'
            ) !== -1
              ? 'text-red-500 hover:text-red-600'
              : 'text-gray-800 hover:text-gray-600'}"
          >
            <i
              class="fas fa-table mr-2 text-sm {location.href.indexOf(
                '/containers'
              ) !== -1
                ? 'opacity-75'
                : 'text-gray-400'}"
            />
            Containers
          </a>
        </li>
        <li class="items-center">
          <a
            use:link
            href="/images"
            class="text-xs uppercase py-3 font-bold block {location.href.indexOf(
              '/images'
            ) !== -1
              ? 'text-red-500 hover:text-red-600'
              : 'text-gray-800 hover:text-gray-600'}"
          >
            <i
              class="fas fa-table mr-2 text-sm {location.href.indexOf(
                '/images'
              ) !== -1
                ? 'opacity-75'
                : 'text-gray-400'}"
            />
            Base Image
          </a>
        </li>
        <li class="items-center">
          <a
            use:link
            href="/tables"
            class="text-xs uppercase py-3 font-bold block {location.href.indexOf(
              '/tables'
            ) !== -1
              ? 'text-red-500 hover:text-red-600'
              : 'text-gray-800 hover:text-gray-600'}"
          >
            <i
              class="fas fa-table mr-2 text-sm {location.href.indexOf(
                '/tables'
              ) !== -1
                ? 'opacity-75'
                : 'text-gray-400'}"
            />
            Tables
          </a>
        </li>

        <li class="items-center">
          <a
            use:link
            href="/settings"
            class="text-xs uppercase py-3 font-bold block {location.href.indexOf(
              '/settings'
            ) !== -1
              ? 'text-red-500 hover:text-red-600'
              : 'text-gray-800 hover:text-gray-600'}"
          >
            <i
              class="fas fa-tools mr-2 text-sm {location.href.indexOf(
                '/settings'
              ) !== -1
                ? 'opacity-75'
                : 'text-gray-400'}"
            />
            Settings
          </a>
        </li>
      </ul>

      <!-- Divider -->
      <hr class="my-4 md:min-w-full" />
      <!-- Heading -->
      <h6
        class="md:min-w-full text-gray-600 text-xs uppercase font-bold block pt-1 pb-4 no-underline"
      >
        Auth Layout Pages
      </h6>
      <!-- Navigation -->

      <ul class="md:flex-col md:min-w-full flex flex-col list-none md:mb-4">
        <li class="items-center">
          <a
            use:link
            class="text-gray-800 hover:text-gray-600 text-xs uppercase py-3 font-bold block"
            href="/login"
          >
            <i class="fas fa-fingerprint text-gray-400 mr-2 text-sm" />
            Login
          </a>
        </li>

        <li class="items-center">
          <a
            use:link
            class="text-gray-800 hover:text-gray-600 text-xs uppercase py-3 font-bold block"
            href="/register"
          >
            <i class="fas fa-clipboard-list text-gray-400 mr-2 text-sm" />
            Register
          </a>
        </li>
      </ul>
    </div>
  </div>
</nav>
