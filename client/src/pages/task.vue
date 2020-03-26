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
                    <button v-if="completed == false" type="button" class="btn btn-primary" 
                    data-toggle="modal" data-target="#modal-add">New Task
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
                          	<form v-model="FormTask">
                          		<div class="form-group">
                          			<label for="inputDescription">Description</label>
                          			<textarea v-model="FormTask.description" class="form-control" rows="4" placeholder="Ex.: Building the Navigation on android UI" />
                          		</div>
                          	</form>
                          </div>
                          <div class="modal-footer justify-content-between">
                            <button type="button" class="btn btn-default" data-dismiss="modal">Close
                            </button>
                            <button type="button" class="btn btn-primary" data-dismiss="modal" 
                              @click="insertTask(FormTask)">Create
                            </button>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <br>
                </div>

                <div v-if="values.length != 0" class="col-md-12">
                	<div class="card">
                		<table class="table table-bordered">
                			<thead>
                				<tr>
                					<th style="width: 10px">#</th>
                					<th>Task</th>
                					<th>Time Spent</th>
                					<th>Actions</th>
                				</tr>
                			</thead>
                			<tbody>
                				<tr v-for="(v , k) in values">
                					<td>{{ (k + 1) }}</td>
                					<td>{{ v.description }}</td>
                					<td>{{ v.time_spent }}h</td>
                					<td>
                						<a class="btn btn-primary" :href="'/projects/' + id + '/' + v.guid">
                							<i class="fas fa-edit"></i> Edit
                						</a>
                						<!--<TaskItem :title.sync="title" :id="v.guid" :description.sync="v.description">
                						</TaskItem>-->
                						<!--<button type="button" class="btn btn-primary" data-toggle="modal" data-target="#modal-edit"> <i class="fas fa-edit"></i> Edit
                						</button>-->
                						<button class="btn btn-danger" @click="deleteTask(v.guid)">
                							<i class="fas fa-eraser"></i> Delete
                						</button>
                					</td>
                				</tr>
                			</tbody>
                		</table>
                	</div>
                </div>
              </div>
            </MainContent>
            <Footbar></Footbar>
          </div>
        </template>
        <template v-else slot="content">
          <div>
            Erro - Servidor nao disponivel. <a :href="'/projects/' + id">Clique</a>
            para Recarregar a Pagina
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
import TaskItem from "components/TaskItem.vue"
export default {
    data(){
      return{
        title: '',
        pid: '',
        tid: '',
        completed: '',
        values: {},
        FormTask: {
          description: ''
        },
        error: false
      }
    },
    created () {
      this.title = 'Tasks'
      axios.get("http://localhost:8080/projects/" + this.id).then((r) => {

        var v = r.data.data
        this.completed = v.completed
        //this.values = v.data
        //console.log(this.values) 
        //console.log(this.values.completed)
      }).catch((e) => {
        this.error = true
        //console.log(e)
      })

      axios.get("http://localhost:8080/tasks/" + this.id).then((r) => {
        var v = r.data
        this.values = v.data
        //console.log(this.values) 
        //console.log(this.values.completed)
      }).catch((e) => {
        this.error = true
        //console.log(e)
      })

      //console.log("http://localhost:8080/projects/" + this.id)
    },
    methods: {
      insertTask(form) {
        let values = {description: form.description, project_id: this.id}
        axios({
          method: 'post',
          url: 'http://localhost:8080/projects',
          headers: {"Content-Type": "application/json", "Accept": "application/json"},
          data: values
        }).then((r) => { 
          //console.log(r)
          this.FormTask.description = ''
          window.location.reload()
          })
      },
      deleteTask(guid) {
          axios({
            method: 'delete',
            url: 'http://localhost:8080/tasks/' + guid,
            headers: {"Content-Type": "application/json", "Accept": "application/json"},
          }).then((r) => { 
            //console.log(r)
            window.location.reload()
          })
          //console.log("id => " + guid)
        }
    },
    components: {
        Parent,
        Navbar,
        Sidebar,
        Footbar,
        MainContent,
        ProjectItem,
        TaskItem
    }
}
</script>
