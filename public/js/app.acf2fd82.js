(function(e){function t(t){for(var n,r,l=t[0],s=t[1],d=t[2],c=0,m=[];c<l.length;c++)r=l[c],Object.prototype.hasOwnProperty.call(i,r)&&i[r]&&m.push(i[r][0]),i[r]=0;for(n in s)Object.prototype.hasOwnProperty.call(s,n)&&(e[n]=s[n]);u&&u(t);while(m.length)m.shift()();return o.push.apply(o,d||[]),a()}function a(){for(var e,t=0;t<o.length;t++){for(var a=o[t],n=!0,r=1;r<a.length;r++){var s=a[r];0!==i[s]&&(n=!1)}n&&(o.splice(t--,1),e=l(l.s=a[0]))}return e}var n={},i={app:0},o=[];function r(e){return l.p+"js/"+({about:"about"}[e]||e)+"."+{about:"a1297950"}[e]+".js"}function l(t){if(n[t])return n[t].exports;var a=n[t]={i:t,l:!1,exports:{}};return e[t].call(a.exports,a,a.exports,l),a.l=!0,a.exports}l.e=function(e){var t=[],a=i[e];if(0!==a)if(a)t.push(a[2]);else{var n=new Promise((function(t,n){a=i[e]=[t,n]}));t.push(a[2]=n);var o,s=document.createElement("script");s.charset="utf-8",s.timeout=120,l.nc&&s.setAttribute("nonce",l.nc),s.src=r(e);var d=new Error;o=function(t){s.onerror=s.onload=null,clearTimeout(c);var a=i[e];if(0!==a){if(a){var n=t&&("load"===t.type?"missing":t.type),o=t&&t.target&&t.target.src;d.message="Loading chunk "+e+" failed.\n("+n+": "+o+")",d.name="ChunkLoadError",d.type=n,d.request=o,a[1](d)}i[e]=void 0}};var c=setTimeout((function(){o({type:"timeout",target:s})}),12e4);s.onerror=s.onload=o,document.head.appendChild(s)}return Promise.all(t)},l.m=e,l.c=n,l.d=function(e,t,a){l.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:a})},l.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},l.t=function(e,t){if(1&t&&(e=l(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var a=Object.create(null);if(l.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)l.d(a,n,function(t){return e[t]}.bind(null,n));return a},l.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return l.d(t,"a",t),t},l.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},l.p="/",l.oe=function(e){throw console.error(e),e};var s=window["webpackJsonp"]=window["webpackJsonp"]||[],d=s.push.bind(s);s.push=t,s=s.slice();for(var c=0;c<s.length;c++)t(s[c]);var u=d;o.push([0,"chunk-vendors"]),a()})({0:function(e,t,a){e.exports=a("56d7")},"56d7":function(e,t,a){"use strict";a.r(t);a("e260"),a("e6cf"),a("cca6"),a("a79d");var n=a("2b0e"),i=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("v-app",{attrs:{id:"app"}},[a("Navbar"),a("router-view")],1)},o=[],r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("nav",[a("v-app-bar",{attrs:{color:"primary"}},[a("v-app-bar-title",[e._v("Go Crud")])],1)],1)},l=[],s=a("2877"),d=a("6544"),c=a.n(d),u=a("40dc"),m=a("bb78"),f={},v=Object(s["a"])(f,r,l,!1,null,null,null),p=v.exports;c()(v,{VAppBar:u["a"],VAppBarTitle:m["a"]});var h={components:{Navbar:p}},b=h,g=a("7496"),x=Object(s["a"])(b,i,o,!1,null,null,null),I=x.exports;c()(x,{VApp:g["a"]});var y=a("f309");n["a"].use(y["a"]);var _=new y["a"]({}),k=(a("d3b7"),a("3ca3"),a("ddb0"),a("8c4f")),w=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("v-container",[a("v-row",[a("v-col",{attrs:{cols:"12",md:"12"}},[a("v-data-table",{staticClass:"elevation-1",attrs:{headers:e.headers,items:e.student,loading:e.loading1},scopedSlots:e._u([{key:"top",fn:function(){return[a("v-toolbar",{attrs:{flat:""}},[a("v-spacer"),a("v-dialog",{attrs:{"max-width":"500px"},scopedSlots:e._u([{key:"activator",fn:function(t){var n=t.on,i=t.attrs;return[a("v-btn",e._g(e._b({staticClass:"mb-2",attrs:{color:"primary",dark:""}},"v-btn",i,!1),n),[e._v(" Add Student ")])]}}]),model:{value:e.dialog,callback:function(t){e.dialog=t},expression:"dialog"}},[a("v-card",[a("v-card-title",[a("span",{staticClass:"text-h6 ml-2"},[e._v(e._s(e.formTitle))])]),a("v-card-text",[a("v-container",[a("v-row",[a("v-col",{attrs:{cols:"12",md:"6"}},[a("v-text-field",{attrs:{label:"First Name",outlined:"",dense:""},model:{value:e.editedItem.firstName,callback:function(t){e.$set(e.editedItem,"firstName",t)},expression:"editedItem.firstName"}})],1),a("v-col",{attrs:{cols:"12",md:"6"}},[a("v-text-field",{attrs:{label:"Last Name",outlined:"",dense:""},model:{value:e.editedItem.lastName,callback:function(t){e.$set(e.editedItem,"lastName",t)},expression:"editedItem.lastName"}})],1),a("v-col",{attrs:{cols:"12",md:"6"}},[a("v-text-field",{attrs:{label:"Email",outlined:"",type:"email",dense:""},model:{value:e.editedItem.email,callback:function(t){e.$set(e.editedItem,"email",t)},expression:"editedItem.email"}})],1),a("v-col",{attrs:{cols:"12",md:"6"}},[a("v-text-field",{attrs:{label:"Age",outlined:"",type:"number",dense:""},model:{value:e.editedItem.age,callback:function(t){e.$set(e.editedItem,"age",t)},expression:"editedItem.age"}})],1),a("v-col",{attrs:{cols:"12",md:"6"}},[a("v-select",{attrs:{items:e.items,outlined:"",dense:"",label:"Gender"},model:{value:e.editedItem.gender,callback:function(t){e.$set(e.editedItem,"gender",t)},expression:"editedItem.gender"}})],1)],1)],1)],1),a("v-card-actions",[a("v-spacer"),a("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:e.close}},[e._v(" Cancel ")]),a("v-btn",{attrs:{color:"blue darken-1",text:"",loading:e.loading2},on:{click:e.save}},[e._v(" Save ")])],1)],1)],1),a("v-dialog",{attrs:{"max-width":"500px"},model:{value:e.dialogDelete,callback:function(t){e.dialogDelete=t},expression:"dialogDelete"}},[a("v-card",[a("v-card-title",{staticClass:"text-h5"},[e._v("Are you sure you want to delete this Student?")]),a("v-card-actions",[a("v-spacer"),a("v-btn",{attrs:{color:"primary",text:""},on:{click:e.closeDelete}},[e._v("Cancel")]),a("v-btn",{attrs:{color:"primary",text:""},on:{click:e.deleteItemConfirm}},[e._v("YES")]),a("v-spacer")],1)],1)],1)],1)]},proxy:!0},{key:"item.actions",fn:function(t){var n=t.item;return[a("v-btn",{staticClass:"mr-2",attrs:{small:"",color:"primary"}},[e._v(" view ")]),a("v-btn",{staticClass:"mr-2",attrs:{small:"",color:"primary"},on:{click:function(t){return e.editItem(n)}}},[e._v(" edit ")]),a("v-btn",{attrs:{small:"",color:"primary"},on:{click:function(t){return e.deleteItem(n)}}},[e._v(" delete ")])]}},{key:"item.name",fn:function(t){var n=t.item;return[a("span",[e._v(e._s(n.firstName)+" "+e._s(n.lastName))])]}}],null,!0)})],1)],1)],1)},O=[],N=(a("a434"),a("e9c4"),a("bc3a")),j=a.n(N),C={name:"Home",components:{},data:function(){return{dialog:!1,loading1:!1,loading2:!1,dialogDelete:!1,items:["male","female"],headers:[{text:"Name",align:"start",sortable:!1,value:"name"},{text:"Email",value:"email",sortable:!1},{text:"Gender",value:"gender",sortable:!1},{text:"Age",value:"age",sortable:!1},{text:"Actions",value:"actions",sortable:!1}],desserts:[],editedIndex:-1,editedItem:{firstName:"",lastName:"",email:"",age:0,gender:"",id:0},student:[]}},computed:{formTitle:function(){return-1===this.editedIndex?"New Student":"Edit Student"}},watch:{dialog:function(e){e||this.close()},dialogDelete:function(e){e||this.closeDelete()}},methods:{fetchStudent:function(){var e=this;this.loading1=!0,j.a.get("http://localhost:3000/student",{headers:{"Content-Type":"application/json"}}).then((function(t){e.loading1=!1,e.student=t.body})).catch((function(e){return console.log(e)}))},editItem:function(e){this.editedIndex=this.student.indexOf(e),this.editedItem=Object.assign({},e),this.dialog=!0},deleteItem:function(e){this.editedIndex=this.student.indexOf(e),this.editedItem=Object.assign({},e),this.dialogDelete=!0},deleteItemConfirm:function(){this.student.splice(this.editedIndex,1),this.closeDelete()},close:function(){var e=this;this.dialog=!1,this.$nextTick((function(){e.editedItem=Object.assign({},e.defaultItem),e.editedIndex=-1}))},closeDelete:function(){var e=this;this.dialogDelete=!1,this.$nextTick((function(){e.editedItem=Object.assign({},e.defaultItem),e.editedIndex=-1}))},save:function(){var e=this;this.editedIndex>-1?Object.assign(this.student[this.editedIndex],this.editedItem):(this.loading2=!0,fetch("http://localhost:3000/student",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({firstName:this.editedItem.firstName,lastName:this.editedItem.lastName,email:this.editedItem.email,age:JSON.parse(this.editedItem.age),gender:this.editedItem.gender,id:0})}).then((function(t){e.loading2=!1,console.log(t),e.student.push(e.editedItem)})).catch((function(t){console.log(t),e.loading2=!1}))),this.close()}},mounted:function(){this.fetchStudent()}},S=C,T=a("8336"),V=a("b0af"),D=a("99d9"),A=a("62ad"),$=a("a523"),E=a("8fea"),P=a("169a"),J=a("0fd9"),M=a("b974"),B=a("2fa4"),G=a("8654"),L=a("71d9"),F=Object(s["a"])(S,w,O,!1,null,null,null),H=F.exports;c()(F,{VBtn:T["a"],VCard:V["a"],VCardActions:D["a"],VCardText:D["b"],VCardTitle:D["c"],VCol:A["a"],VContainer:$["a"],VDataTable:E["a"],VDialog:P["a"],VRow:J["a"],VSelect:M["a"],VSpacer:B["a"],VTextField:G["a"],VToolbar:L["a"]}),n["a"].use(k["a"]);var q=[{path:"/",name:"Home",component:H},{path:"/about",name:"About",component:function(){return a.e("about").then(a.bind(null,"f820"))}}],R=new k["a"]({routes:q}),Y=R;n["a"].config.productionTip=!1,new n["a"]({vuetify:_,router:Y,render:function(e){return e(I)}}).$mount("#app")}});
//# sourceMappingURL=app.acf2fd82.js.map