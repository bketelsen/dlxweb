<script>
  // core components
  import CardContainers from "components/Cards/CardContainers.svelte";

  import CardBarChart from "components/Cards/CardBarChart.svelte";
  import CardPageVisits from "components/Cards/CardPageVisits.svelte";
  import CardSocialTraffic from "components/Cards/CardSocialTraffic.svelte";
  export let location;
</script>

<div>
  <div class="flex flex-wrap">
    <div class="w-full px-4 mb-12 xl:w-8/12 xl:mb-0">
      <CardContainers />
    </div>
    <div class="w-full px-4 xl:w-4/12">
      <CardBarChart />
    </div>
  </div>
  <div class="flex flex-wrap mt-4">
    <div class="w-full px-4 mb-12 xl:w-8/12 xl:mb-0">
      <CardPageVisits />
    </div>
    <div class="w-full px-4 xl:w-4/12">
      <CardSocialTraffic />
    </div>
  </div>
</div>
