<template>
    <parent>
        <template slot="title">Devlogbook | {{ title }}</template>
        <template v-if="error == false" slot="content">
          <div class="wrapper">
            <Navbar></Navbar>
            <Sidebar>
              <ul slot="menu-item" class="nav nav-pills nav-sidebar flex-column"
              data-widget="treeview" role="menu" data-accordion="false">
                <li class="nav-item">
                  <a href="/" class="nav-link active">
                    <i class="nav-icon fas fa-th"></i>
                    <p>Dashboard</p>
                  </a>
                </li>
                <li class="nav-item">
                  <a href="/projects" class="nav-link">
                    <!--<i class="nav-icon fas fa-copy"></i>-->
                    <i class="nav-icon fas fa-boxes"></i>
                    <p>Projects</p>
                  </a>
                </li>
              </ul>
            </Sidebar>
            <MainContent>
              <h1 slot="text-1" class="m-0 text-dark">{{ title }}</h1>
              <ol slot="text-2" class="breadcrumb float-sm-right">
                <li class="breadcrumb-item"><a href="/">Home</a></li>
                <li class="breadcrumb-item active">{{ title }}</li>
              </ol>
              <div slot="inner-app" class="row">
                <div class="col-lg-6 col-12">
                  <!-- small box -->
                  <div class="small-box bg-info">
                    <div class="inner">
                      <!--<h1><%= data.status.total %></h1>-->
                      <StatusItem :value="values.total" :description="'Total Projects'">
                      </StatusItem>
                    </div>
                    <div class="icon">
                      <i class="fas fa-boxes"></i>
                    </div>
                  </div>
                </div>
                <!-- ./col -->

                <div class="col-lg-6 col-12">
                  <!-- small box -->
                  <div class="small-box bg-danger">
                    <div class="inner">
                      <!--<h1><%= data.status.total %></h1>-->
                      <StatusItem :value="values.completed" :description="'Completed Projects'">
                      </StatusItem>
                    </div>
                    <div class="icon">
                      <i class="fas fa-box"></i>
                    </div>
                  </div>
                </div>
                <!-- ./col -->

                <div class="col-lg-6 col-12">
                  <!-- small box -->
                  <div class="small-box bg-warning">
                    <div class="inner">
                      <!--<h1><%= data.status.total %></h1>-->
                      <StatusItem :value="values.average_time + 'h'" :description="'Average Time Spent p/ Project'">
                      </StatusItem>
                    </div>
                    <div class="icon">
                      <i class="fas fa-clock"></i>
                    </div>
                  </div>
                </div>
                <!-- ./col -->

              </div>
            </MainContent>
            <Footbar></Footbar>
          </div>
        </template>
        <template v-else slot="content">
          <div>
            Erro - Servidor nao disponivel. <a href="/">Clique</a> para Recarregar a Pagina
          </div>
        </template>
    </parent>
</template>
<script>
import axios from "axios"
import Parent from "common/parent.vue"
import Navbar from "common/navbar.vue"
import Sidebar from "common/sidebar.vue"
import Footbar from "common/footbar.vue"
import MainContent from "common/main_content.vue"
import StatusItem from "components/StatusItem.vue"

export default {
    data (){
      return{
        title: '',
        values: {},
        error: false
      }
    },
    created () {
      this.title = 'Dashboard'
      axios.get('http://localhost:8080/status').then((r) => {
        var v = r.data
        this.values = v.data
        //console.log(this.values.completed)
      }).catch((e) => {
        this.error = true
        //console.log(e)
      })
    },
    methods: {
    },
    components: {
        Parent,
        Navbar,
        Sidebar,
        Footbar,
        MainContent,
        StatusItem
    }
}
</script>
