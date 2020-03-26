<template>
  <!-- Content Wrapper. Contains page content -->
  <div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <div class="content-header">
      <div class="container-fluid">
        <div class="row mb-2">
          <div class="col-sm-6">
            <slot name="text-1"></slot>
          </div><!-- /.col -->
          <div class="col-sm-6">
            <slot name="text-2"></slot>
          </div><!-- /.col -->
        </div><!-- /.row -->
      </div><!-- /.container-fluid -->
    </div>
    <!-- /.content-header -->
    <!-- Main content -->
    <section class="content">
      <div class="container-fluid">
        <!-- Small boxes (Stat box) -->
        <slot name="inner-app"></slot>
      </div>
    </section>
  </div>
</template>
<script>
  export default {
    name: 'maincontent'
  }
</script>
