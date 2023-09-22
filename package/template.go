
package main

type TemplateMetadata struct {
  Name string
  Title string
  Description string
  Owner string
}

type TemplateSpec struct {
  Owner string
  Type string
  Resource *Resource
}

type Resource struct {
  Name string
  Parameters []Parameter
  DeletionPolicy string
}

type Parameter struct {

}

