<!-- App.svelte -->
<script>
  import { Router, Route } from "svelte-routing";

  // Admin Layout
  import Admin from "./layouts/Admin.svelte";
  // Auth Layout
  import Auth from "./layouts/Auth.svelte";


  // Toast
  import { SvelteToast } from "@zerodevx/svelte-toast";
  const options = {
    duration: 4000, // duration of progress bar tween to the `next` value
    initial: 1, // initial progress bar value
    next: 0, // next progress value
    pausable: false, // pause progress bar tween on mouse hover
    dismissable: true, // allow dismiss with close button
    reversed: false, // insert new toast to bottom of stack
    intro: { x: 256 }, // toast intro fly animation settings
    theme: {}, // css var overrides
  };

  export let url = "";
</script>

<SvelteToast {options} />

<Router {url}>
  <!-- auth layout -->
  <Route path="auth/*auth" component={Auth} />

  <!-- admin layout -->
  <Route path="/*" component={Admin} />
</Router>
