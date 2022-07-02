<template>
  <div class="py-12 bg-white">
    <div>Snack Machine</div>
    <div class="flex flex-row justify-center">
      <div class="basis-1/2 flex justify-center">
        <div class="flex flex-col">
          <div class="flex flex-col mt-4">
            <div class="flex flex-row">
              <div class="h-20 w-20">
                <img src="https://static.bigc.co.th/media/catalog/product/8/8/8858998581153_2.jpg" />
              </div>
              <div>
                <div>$1.00</div>
                <div>10 left</div>
              </div>
            </div>
            <div>
              <button type="button" class="w-full px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 ">
                Buy
              </button>
            </div>
          </div>

          <div class="flex flex-col mt-4">
            <div class="flex flex-row">
              <div class="h-20 w-20">
                <img src="https://static.bigc.co.th/media/catalog/product/8/8/8855199143180_2.jpg" />
              </div>
              <div>
                <div>$1.00</div>
                <div>10 left</div>
              </div>
            </div>
            <div>
              <button type="button" class="w-full px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 ">
                Buy
              </button>
            </div>
          </div>

          <div class="flex flex-col mt-4">
            <div class="flex flex-row">
              <div class="h-20 w-20">
                <img src="https://static.bigc.co.th/media/catalog/product/8/8/8855199143180_2.jpg" />
              </div>
              <div>
                <div>$1.00</div>
                <div>10 left</div>
              </div>
            </div>
            <div>
              <button type="button" class="w-full px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 ">
                Buy
              </button>
            </div>
          </div>
        </div>
      </div>

      <div class="basis-1/2 flex justify-center">
        <div class="flex flex-col">
          <div class="flex flex-row">
            <div class="flex flex-col">
              <div>Money Inserted: $0</div>
              <div class="flex flex-row">
                <button type="button" class="w-full px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 ">
                  Put 1¢
                </button>

                <button type="button" class="ml-2 w-full px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 ">
                  Put 10¢
                </button>

                <button type="button" class="ml-2 w-full px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 ">
                  Put 25¢
                </button>
              </div>
              <div class="flex flex-row mt-4">
                <button type="button" class="w-full px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 ">
                  Put 1$
                </button>

                <button type="button" class="ml-2 w-full px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 ">
                  Put 5$
                </button>

                <button type="button" class="ml-2 w-full px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 ">
                  Put 20$
                </button>
              </div>
              <div class="flex flex-row mt-4">
                <button type="button" class="w-full px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 ">
                  Return
                </button>
              </div>
            </div>
          </div>
          <div class="flex flex-col mt-6">
            <div>Money Inside $0</div>
            <div class="flex flex-row justify-center">
              <div>1¢: 0</div>
              <div>10¢: 0</div>
              <div>25¢: 0</div>
            </div>
            <div class="flex flex-row justify-center">
              <div>1$: 0</div>
              <div>5$: 0</div>
              <div>20$: 0</div>
            </div>
          </div>
        </div>
        
      </div>
    </div>
  </div>
</template>