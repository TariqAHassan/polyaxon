// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon sdk
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.14.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * V1SearchDefinition
 */

public class V1SearchDefinition {
  @SerializedName("query")
  private String query = null;

  @SerializedName("sort")
  private String sort = null;

  @SerializedName("groupby")
  private String groupby = null;

  @SerializedName("columns")
  private String columns = null;

  public V1SearchDefinition query(String query) {
    this.query = query;
    return this;
  }

   /**
   * Get query
   * @return query
  **/
  @ApiModelProperty(value = "")
  public String getQuery() {
    return query;
  }

  public void setQuery(String query) {
    this.query = query;
  }

  public V1SearchDefinition sort(String sort) {
    this.sort = sort;
    return this;
  }

   /**
   * Get sort
   * @return sort
  **/
  @ApiModelProperty(value = "")
  public String getSort() {
    return sort;
  }

  public void setSort(String sort) {
    this.sort = sort;
  }

  public V1SearchDefinition groupby(String groupby) {
    this.groupby = groupby;
    return this;
  }

   /**
   * Get groupby
   * @return groupby
  **/
  @ApiModelProperty(value = "")
  public String getGroupby() {
    return groupby;
  }

  public void setGroupby(String groupby) {
    this.groupby = groupby;
  }

  public V1SearchDefinition columns(String columns) {
    this.columns = columns;
    return this;
  }

   /**
   * Get columns
   * @return columns
  **/
  @ApiModelProperty(value = "")
  public String getColumns() {
    return columns;
  }

  public void setColumns(String columns) {
    this.columns = columns;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1SearchDefinition v1SearchDefinition = (V1SearchDefinition) o;
    return Objects.equals(this.query, v1SearchDefinition.query) &&
        Objects.equals(this.sort, v1SearchDefinition.sort) &&
        Objects.equals(this.groupby, v1SearchDefinition.groupby) &&
        Objects.equals(this.columns, v1SearchDefinition.columns);
  }

  @Override
  public int hashCode() {
    return Objects.hash(query, sort, groupby, columns);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1SearchDefinition {\n");
    
    sb.append("    query: ").append(toIndentedString(query)).append("\n");
    sb.append("    sort: ").append(toIndentedString(sort)).append("\n");
    sb.append("    groupby: ").append(toIndentedString(groupby)).append("\n");
    sb.append("    columns: ").append(toIndentedString(columns)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

