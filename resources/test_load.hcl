name= "test"

includes= "{{ eq '1' '1' }}"

tasks= {
   "judo.echo"= "el nuevo path es {{ .base_path }}"
}