<template>
  <q-page class="row items-stretch bg-light q-pa-sm">
    <div class="col-4 q-pa-sm">
      <q-card
        class="q-pa-sm"
        style="height: 100%"
        color="white"
      >
        <q-card-main>
          <q-btn
            class="full-width"
            label="Get"
            :color="color1"
            @click="Get"
          />
          <q-btn
            class="full-width q-mt-md"
            label="Get by ID"
            :color="color2"
            @click="InputGetById"
          />
          <q-card-actions
            v-if="inputGet"
            class="row q-px-none q-pt-sm"
            align="between"
          >
            <div class="col-10 q-pr-sm">
              <q-input
                class="box no-shadow"
                v-model="id"
                type="number"
                color="white"
                inverted-light
                placeholder="ID"
                @keyup.enter="GetById(id)"
              />
            </div>
            <q-btn
              class="col-2 box"
              outline
              label="Go"
              color="faded"
              @click="GetById(id)"
            />
          </q-card-actions>
          <q-btn
            class="full-width q-mt-md"
            label="Post"
            :color="color3"
            @click="InputPost"
          />
          <q-card-actions
            v-if="inputPost"
            class="q-px-none q-pt-sm"
            align="between"
          >
            <div class="row">
              <div class="col-3 q-pr-xs">
                <q-input
                  class="box no-shadow"
                  v-model="form.id"
                  type="number"
                  color="white"
                  inverted-light
                  placeholder="ID"
                />
              </div>
              <div class="col-9 q-pl-xs">
                <q-input
                  class="box no-shadow"
                  v-model="form.nome"
                  color="white"
                  inverted-light
                  placeholder="Nome"
                />
              </div>
              <div class="col-5 q-pr-xs q-pt-sm">
                <q-input
                  class="box no-shadow"
                  v-model="form.telefone"
                  type="tel"
                  color="white"
                  inverted-light
                  placeholder="Telefone"
                />
              </div>
              <div class="col-7 q-pl-xs q-pt-sm">
                <q-input
                  class="box no-shadow"
                  v-model="form.email"
                  type="email"
                  color="white"
                  inverted-light
                  placeholder="E-mail"
                />
              </div>
              <div class="col-9 q-pr-xs q-pt-sm">
                <q-input
                  class="box no-shadow"
                  v-model="form.endereco.logradouro"
                  color="white"
                  inverted-light
                  placeholder="Logradouro"
                />
              </div>
              <div class="col-3 q-pl-xs q-pt-sm">
                <q-input
                  class="box no-shadow"
                  v-model="form.endereco.numero"
                  type="number"
                  color="white"
                  inverted-light
                  placeholder="Número"
                />
              </div>
              <div class="col-6 q-pr-xs q-pt-sm">
                <q-input
                  class="box no-shadow"
                  v-model="form.endereco.bairro"
                  color="white"
                  inverted-light
                  placeholder="Bairro"
                />
              </div>
              <div class="col-6 q-pl-xs q-pt-sm">
                <q-input
                  class="box no-shadow"
                  v-model="form.endereco.cidade"
                  color="white"
                  inverted-light
                  placeholder="Cidade"
                />
              </div>
              <div class="col-6 q-pr-xs q-pt-sm">
                <q-input
                  class="box no-shadow"
                  v-model="form.endereco.estado"
                  color="white"
                  inverted-light
                  placeholder="Estado"
                  @keyup.enter="Post(form)"
                />
              </div>
              <div class="col-6 q-pl-xs q-pt-sm">
                <q-btn
                  class="box full-width"
                  outline
                  label="Go"
                  color="faded"
                  @click="Post(form)"
                />
              </div>
            </div>
          </q-card-actions>
          <q-btn
            class="full-width q-mt-md"
            label="Delete"
            :color="color4"
            @click="InputDel"
          />
          <q-card-actions
            v-if="inputDel"
            class="row q-px-none q-pt-sm"
            align="between"
          >
            <div class="col-10 q-pr-sm">
              <q-input
                class="box no-shadow"
                v-model="id"
                type="number"
                color="white"
                inverted-light
                placeholder="ID"
                @keyup.enter="Delete(id)"
              />
            </div>
            <q-btn
              class="col-2 box"
              outline
              label="Go"
              color="faded"
              @click="Delete(id)"
            />
          </q-card-actions>
        </q-card-main>
      </q-card>
    </div>
    <div class="col-8 q-pa-sm">
      <q-card
        class="q-pa-lg"
        style="height: 100%"
        color="white"
      >
        <q-card
          v-if="data.length"
          color="white"
        >
          <q-card-title class="bg-faded q-py-sm">
            <strong>{{label}}</strong>
          </q-card-title>
          <q-table
            :data="data"
            :columns="columns"
            row-key="id"
            :loading="loading"
            loading-label="Carregando..."
            :pagination.sync="pagination"
            hide-bottom
            color="faded"
            dense
            grid
          />
        </q-card>
        <q-card
          v-else
          class="row flex-center fit"
          flat
        >
          <q-icon
            name="far fa-folder-open"
            color="faded"
            size="300px"
          />
        </q-card>
      </q-card>
    </div>
  </q-page>
</template>

<script>
const modelForm = {
  id: '',
  nome: '',
  telefone: '',
  email: '',
  endereco: {
    logradouro: '',
    numero: '',
    bairro: '',
    cidade: '',
    estado: ''
  }
}
export default {
  name: 'PageIndex',
  data () {
    return {
      API: process.env.API,
      id: '',
      label: '',
      form: JSON.parse(JSON.stringify(modelForm)),
      data: [],
      columns: [
        { name: 'id', field: 'id', label: 'ID', align: 'left' },
        { name: 'nome', field: 'nome', label: 'Nome', align: 'left' },
        { name: 'telefone', field: 'telefone', label: 'Telefone', align: 'left' },
        { name: 'email', field: 'email', label: 'Email', align: 'left' },
        { name: 'endereco', field: 'endereco', label: 'Endereço', align: 'left' }
      ],
      loading: true,
      pagination: {
        rowsPerPage: 100,
        page: 1
      },
      inputGet: false,
      inputPost: false,
      inputDel: false,
      color1: 'faded',
      color2: 'faded',
      color3: 'faded',
      color4: 'faded'
    }
  },
  methods: {
    InputGetById () {
      this.inputPost = false
      this.inputDel = false
      this.inputGet = !this.inputGet
      if (this.inputGet) {
        this.color1 = 'light'
        this.color2 = 'faded'
        this.color3 = 'light'
        this.color4 = 'light'
      } else {
        this.color1 = 'faded'
        this.color2 = 'faded'
        this.color3 = 'faded'
        this.color4 = 'faded'
      }
    },
    InputPost () {
      this.inputGet = false
      this.inputDel = false
      this.inputPost = !this.inputPost
      if (this.inputPost) {
        this.color1 = 'light'
        this.color2 = 'light'
        this.color3 = 'faded'
        this.color4 = 'light'
      } else {
        this.color1 = 'faded'
        this.color2 = 'faded'
        this.color3 = 'faded'
        this.color4 = 'faded'
      }
    },
    InputDel () {
      this.inputGet = false
      this.inputPost = false
      this.inputDel = !this.inputDel
      if (this.inputDel) {
        this.color1 = 'light'
        this.color2 = 'light'
        this.color3 = 'light'
        this.color4 = 'faded'
      } else {
        this.color1 = 'faded'
        this.color2 = 'faded'
        this.color3 = 'faded'
        this.color4 = 'faded'
      }
    },
    Get () {
      this.color1 = 'faded'
      this.color2 = 'faded'
      this.color3 = 'faded'
      this.color4 = 'faded'
      this.inputGet = false
      this.inputPost = false
      this.inputDel = false
      this.id = ''
      this.form = JSON.parse(JSON.stringify(modelForm))
      this.$axios.get(`${this.API}/cadastros`)
        .then(Res => {
          this.label = 'Get'
          Res.data.map(usuario => {
            usuario.endereco = `${usuario.endereco.logradouro}, ${usuario.endereco.numero} - ${usuario.endereco.bairro} - ${usuario.endereco.cidade} / ${usuario.endereco.estado}`
          })
          this.data = Res.data
        })
        .catch(Err => {
          console.error(Err)
        })
        .finally(() => {
          this.loading = false
        })
    },
    GetById (id) {
      if (id) {
        this.data = []
        this.$axios.get(`${this.API}/cadastro/${id}`)
          .then(Res => {
            this.label = 'Get by ID'
            Res.data.endereco = `${Res.data.endereco.logradouro}, ${Res.data.endereco.numero} - ${Res.data.endereco.bairro} - ${Res.data.endereco.cidade} / ${Res.data.endereco.estado}`
            this.data.push(Res.data)
          })
          .catch(Err => {
            console.error(Err)
          })
          .finally(() => {
            this.loading = false
          })
      }
    },
    Post (form) {
      if (form.id) {
        let usuario = JSON.parse(JSON.stringify(form))
        usuario.endereco.numero = String(usuario.endereco.numero)
        this.$axios.post(`${this.API}/cadastro/${usuario.id}`, usuario)
          .then((Res) => {
            this.$q.notify({
              type: 'positive',
              timeout: 1000,
              message: 'Usuário salvo com sucesso!'
            })
          })
          .catch(Err => {
            console.error(Err)
          })
          .finally(() => {
            this.Get()
          })
      }
    },
    Delete (id) {
      if (id) {
        this.$axios.delete(`${this.API}/cadastro/${id}`)
          .then(() => {
            this.$q.notify({
              type: 'positive',
              timeout: 1000,
              message: 'Usuário excluído com sucesso!'
            })
          })
          .catch(Err => {
            console.error(Err)
          })
          .finally(() => {
            this.Get()
          })
      }
    }
  }
}
</script>
