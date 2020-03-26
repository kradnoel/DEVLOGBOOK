<template>
  <servue>
    <template slot="head">
      <!-- Font Awesome -->
      <link rel="stylesheet" href="/stylesheets/all.min.css">
      <!-- Pace -->
      <link rel="stylesheet" href="/stylesheets/pace-theme-flat-top.css">
      <!-- Theme style -->
      <link rel="stylesheet" href="/stylesheets/adminlte.min.css">
      <link rel="icon" type="image/x-icon" href="/images/favicon.ico"/>
      <title><slot name="title"/></title>
    </template>
    <template slot="content">
      <slot name="content"/>
      <!-- jQuery -->
      <script src="/javascript/jquery.min.js"></script>
      <!-- Popper.js -->
      <!--<script src="/javascript/popper.min.js"></script> -->
      <!-- Bootstrap -->
      <script src="/javascript/bootstrap.bundle.min.js"></script>
      <!-- Pace -->
      <script src="/javascript/pace.min.js"></script>
      <!-- AdminLTE App -->
      <script src="/javascript/adminlte.min.js"></script>
    </template>
  </servue>
</template>
