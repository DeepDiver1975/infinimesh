import { shallowMount } from "@vue/test-utils";
import DeviceManagement from "@/views/DeviceManagement.vue";

describe("DeviceManagement.vue", () => {
  const wrapper = shallowMount(DeviceManagement);

  it("contains radio labels", () => {
    expect(wrapper.vm.radioLabels).toExist();
  });
});
