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
                  <a href="/" class="nav-link">
                    <i class="nav-icon fas fa-th"></i>
                    <p>Dashboard</p>
                  </a>
                </li>
                <li class="nav-item">
                  <a href="/projects" class="nav-link active">
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
                <div class="col-md-12">
                  <div class="card-header">
                    <button type="button" class="btn btn-primary" 
                    data-toggle="modal" data-target="#modal-add">New Project
                    </button>
                    <div class="modal fade" id="modal-add">
                      <div class="modal-dialog">
                        <div class="modal-content">
                          <div class="modal-header">
                            <h4 class="modal-title">{{ title }}</h4>
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                              <span aria-hidden="true">×</span>
                            </button>
                          </div>
                          <div class="modal-body">
                            <form v-model="FormProject">
                              <div class="form-group">
                                <label for="inputName">Project Name</label>
                                <input type="text" v-model="FormProject.title" class="form-control" placeholder="Ex.: Lazz'oz">
                              </div>
                              <div class="form-group">
                                <label for="inputDescription">Project Description</label>
                                <textarea v-model="FormProject.description" class="form-control" rows="4" placeholder="Ex.: Ride sharing app for the 21st century" />
                              </div>
                            </form>
                          </div>
                          <div class="modal-footer justify-content-between">
                            <button type="button" class="btn btn-default" data-dismiss="modal">Close
                            </button>
                            <button type="button" class="btn btn-primary" data-dismiss="modal" 
                              @click="insertProject(FormProject)">Create
                            </button>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <br>
                </div>

                <div v-for="value in values" class="col-lg-6 col-md-12">
                  <div v-if="value.completed == false" class="small-box bg-success">
                    <div class="" style="float: right; padding: 0.5em;">
                      <button type="button" class="btn btn-success dropup dropdown-hover dropdown-icon dropdown-hover" data-toggle="dropdown"><i class="fas fa-bars"></i>
                      </button>
                      <div class="dropdown-menu dropdown-menu-right" role="menu">
                        <a class="dropdown-item" :href="'/projects/' + value.guid">
                          <i class="fas fa-edit"></i> Edit
                        </a>
                        <div class="dropdown-divider"></div>
                        <button type="button" class="dropdown-item" 
                        style="color:red; font-family: arial;" @click="deleteProject(value.guid)">
                          <i class="fas fa-eraser"></i> Delete
                        </button>
                      </div>
                    </div>
                    <div class="inner">
                      <ProjectItem :title="value.title" :description="value.description" 
                        :timespent="value.time_spent">
                      </ProjectItem>
                    </div>
                  </div>
                  <div v-else class="small-box bg-danger">
                    <div class="" style="float: right; padding: 0.5em;">
                      <button type="button" class="btn btn-danger dropup dropdown-hover dropdown-icon dropdown-hover" data-toggle="dropdown"><i class="fas fa-bars"></i>
                      </button>
                      <div class="dropdown-menu dropdown-menu-right" role="menu">
                        <a class="dropdown-item" :href="'/projects/' + value.guid">
                          <i class="fas fa-edit"></i> Edit
                        </a>
                        <div class="dropdown-divider"></div>
                        <button type="button" class="dropdown-item" 
                        style="color:red; font-family: arial;" @click="deleteProject(value.guid)">
                          <i class="fas fa-eraser"></i> Delete
                        </button>
                      </div>
                    </div>
                    <div class="inner">
                      <ProjectItem :title="value.title" :description="value.description" 
                        :timespent="value.time_spent">
                      </ProjectItem>
                    </div>
                  </div>
                </div>
              </div>
            </MainContent>
            <Footbar></Footbar>
          </div>
        </template>
        <template v-else slot="content">
          <div>
            Erro - Servidor nao disponivel. <a href="/projects">Clique</a> para Recarregar a Pagina
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
import ProjectItem from "components/ProjectItem.vue"
export default {
    data(){
      return{
        title: '',
        values: {},
        FormProject: {
          title: '',
          description: '',
        },
        error: false
      }
    },
    created () {
      this.title = 'Projects'
      axios.get('http://localhost:8080/projects').then((r) => {
        var v = r.data
        this.values = v.data
        //console.log(this.values) 
        //console.log(this.values.completed)
      }).catch((e) => {
        this.error = true
        //console.log(e)
      })
    },
    methods: {
      insertProject(form) {
        let values = {title: form.title, description: form.description}
        axios({
          method: 'post',
          url: 'http://localhost:8080/projects',
          headers: {"Content-Type": "application/json", "Accept": "application/json"},
          data: values
        }).then((r) => { 
          //console.log(r)
          this.FormProject.title = ''
          this.FormProject.description = ''
          window.location.reload()
          })
      },
      deleteProject(guid) {
        axios({
          method: 'delete',
          url: 'http://localhost:8080/projects/' + guid,
          headers: {"Content-Type": "application/json", "Accept": "application/json"},
        }).then((r) => { 
          //console.log(r)
          window.location.reload()
        })
      }
    },
    components: {
        Parent,
        Navbar,
        Sidebar,
        Footbar,
        MainContent,
        ProjectItem
    }
}
</script>
