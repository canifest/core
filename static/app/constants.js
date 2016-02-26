angular.module('canifestApp.constants', [])
  /**
  * Sets the constants required for the router service
  */
  .constant(
    'ROUTES',
    {
      /**
       * Operations page
       */
      OPERATION_STATE: 'operations',
      OPERATION_STATE_URL:'/operations'
    }
  );
