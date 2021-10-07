<script>
  // can be one of light or dark
  import { onMount } from "svelte";
  import { ProjectService } from '../../oto';
  import ProjectDropdown  from 'components/Dropdowns/ProjectDropdown.svelte'
  let projects = [];

  onMount(async () => {
    const projectService = new ProjectService();
    const result = await projectService.list();
    projects = result.projects
    console.log(projects)

  });

</script>
<!-- Navbar -->
<nav
  class="absolute top-0 left-0 w-full z-10 bg-transparent md:flex-row md:flex-no-wrap md:justify-start flex items-center p-4"
>
  <div
    class="w-full mx-autp items-center flex justify-between md:flex-no-wrap flex-wrap md:px-10 px-4"
  >
    <!-- Brand -->
    <a
      class="text-white text-sm uppercase hidden lg:inline-block font-semibold"
      href="#dashboard" on:click={(e) => e.preventDefault()}
    >
      Dashboard
    </a>
    <!-- Form -->
    <form
      class="md:flex hidden flex-row flex-wrap items-center lg:ml-auto mr-3"
    >
      <div class="relative flex w-full flex-wrap items-stretch">
        <ProjectDropdown projects={projects}/>
      </div>
    </form>

  </div>
</nav>
<!-- End Navbar -->
