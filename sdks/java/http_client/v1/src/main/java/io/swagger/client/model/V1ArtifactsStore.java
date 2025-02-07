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
import java.util.ArrayList;
import java.util.List;
import org.threeten.bp.OffsetDateTime;

/**
 * V1ArtifactsStore
 */

public class V1ArtifactsStore {
  @SerializedName("uuid")
  private String uuid = null;

  @SerializedName("name")
  private String name = null;

  @SerializedName("description")
  private String description = null;

  @SerializedName("readme")
  private String readme = null;

  @SerializedName("tags")
  private List<String> tags = null;

  @SerializedName("created_at")
  private OffsetDateTime createdAt = null;

  @SerializedName("updated_at")
  private OffsetDateTime updatedAt = null;

  @SerializedName("frozen")
  private Boolean frozen = null;

  @SerializedName("disabled")
  private Boolean disabled = null;

  @SerializedName("deleted")
  private Boolean deleted = null;

  @SerializedName("k8s_secret")
  private String k8sSecret = null;

  @SerializedName("type")
  private String type = null;

  @SerializedName("mount_path")
  private String mountPath = null;

  @SerializedName("host_path")
  private String hostPath = null;

  @SerializedName("volume_claim")
  private String volumeClaim = null;

  @SerializedName("bucket")
  private String bucket = null;

  @SerializedName("read_only")
  private Boolean readOnly = null;

  public V1ArtifactsStore uuid(String uuid) {
    this.uuid = uuid;
    return this;
  }

   /**
   * Get uuid
   * @return uuid
  **/
  @ApiModelProperty(value = "")
  public String getUuid() {
    return uuid;
  }

  public void setUuid(String uuid) {
    this.uuid = uuid;
  }

  public V1ArtifactsStore name(String name) {
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @ApiModelProperty(value = "")
  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public V1ArtifactsStore description(String description) {
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @ApiModelProperty(value = "")
  public String getDescription() {
    return description;
  }

  public void setDescription(String description) {
    this.description = description;
  }

  public V1ArtifactsStore readme(String readme) {
    this.readme = readme;
    return this;
  }

   /**
   * Get readme
   * @return readme
  **/
  @ApiModelProperty(value = "")
  public String getReadme() {
    return readme;
  }

  public void setReadme(String readme) {
    this.readme = readme;
  }

  public V1ArtifactsStore tags(List<String> tags) {
    this.tags = tags;
    return this;
  }

  public V1ArtifactsStore addTagsItem(String tagsItem) {
    if (this.tags == null) {
      this.tags = new ArrayList<String>();
    }
    this.tags.add(tagsItem);
    return this;
  }

   /**
   * Get tags
   * @return tags
  **/
  @ApiModelProperty(value = "")
  public List<String> getTags() {
    return tags;
  }

  public void setTags(List<String> tags) {
    this.tags = tags;
  }

  public V1ArtifactsStore createdAt(OffsetDateTime createdAt) {
    this.createdAt = createdAt;
    return this;
  }

   /**
   * Get createdAt
   * @return createdAt
  **/
  @ApiModelProperty(value = "")
  public OffsetDateTime getCreatedAt() {
    return createdAt;
  }

  public void setCreatedAt(OffsetDateTime createdAt) {
    this.createdAt = createdAt;
  }

  public V1ArtifactsStore updatedAt(OffsetDateTime updatedAt) {
    this.updatedAt = updatedAt;
    return this;
  }

   /**
   * Get updatedAt
   * @return updatedAt
  **/
  @ApiModelProperty(value = "")
  public OffsetDateTime getUpdatedAt() {
    return updatedAt;
  }

  public void setUpdatedAt(OffsetDateTime updatedAt) {
    this.updatedAt = updatedAt;
  }

  public V1ArtifactsStore frozen(Boolean frozen) {
    this.frozen = frozen;
    return this;
  }

   /**
   * Get frozen
   * @return frozen
  **/
  @ApiModelProperty(value = "")
  public Boolean isFrozen() {
    return frozen;
  }

  public void setFrozen(Boolean frozen) {
    this.frozen = frozen;
  }

  public V1ArtifactsStore disabled(Boolean disabled) {
    this.disabled = disabled;
    return this;
  }

   /**
   * Get disabled
   * @return disabled
  **/
  @ApiModelProperty(value = "")
  public Boolean isDisabled() {
    return disabled;
  }

  public void setDisabled(Boolean disabled) {
    this.disabled = disabled;
  }

  public V1ArtifactsStore deleted(Boolean deleted) {
    this.deleted = deleted;
    return this;
  }

   /**
   * Get deleted
   * @return deleted
  **/
  @ApiModelProperty(value = "")
  public Boolean isDeleted() {
    return deleted;
  }

  public void setDeleted(Boolean deleted) {
    this.deleted = deleted;
  }

  public V1ArtifactsStore k8sSecret(String k8sSecret) {
    this.k8sSecret = k8sSecret;
    return this;
  }

   /**
   * Get k8sSecret
   * @return k8sSecret
  **/
  @ApiModelProperty(value = "")
  public String getK8sSecret() {
    return k8sSecret;
  }

  public void setK8sSecret(String k8sSecret) {
    this.k8sSecret = k8sSecret;
  }

  public V1ArtifactsStore type(String type) {
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @ApiModelProperty(value = "")
  public String getType() {
    return type;
  }

  public void setType(String type) {
    this.type = type;
  }

  public V1ArtifactsStore mountPath(String mountPath) {
    this.mountPath = mountPath;
    return this;
  }

   /**
   * Get mountPath
   * @return mountPath
  **/
  @ApiModelProperty(value = "")
  public String getMountPath() {
    return mountPath;
  }

  public void setMountPath(String mountPath) {
    this.mountPath = mountPath;
  }

  public V1ArtifactsStore hostPath(String hostPath) {
    this.hostPath = hostPath;
    return this;
  }

   /**
   * Get hostPath
   * @return hostPath
  **/
  @ApiModelProperty(value = "")
  public String getHostPath() {
    return hostPath;
  }

  public void setHostPath(String hostPath) {
    this.hostPath = hostPath;
  }

  public V1ArtifactsStore volumeClaim(String volumeClaim) {
    this.volumeClaim = volumeClaim;
    return this;
  }

   /**
   * Get volumeClaim
   * @return volumeClaim
  **/
  @ApiModelProperty(value = "")
  public String getVolumeClaim() {
    return volumeClaim;
  }

  public void setVolumeClaim(String volumeClaim) {
    this.volumeClaim = volumeClaim;
  }

  public V1ArtifactsStore bucket(String bucket) {
    this.bucket = bucket;
    return this;
  }

   /**
   * Get bucket
   * @return bucket
  **/
  @ApiModelProperty(value = "")
  public String getBucket() {
    return bucket;
  }

  public void setBucket(String bucket) {
    this.bucket = bucket;
  }

  public V1ArtifactsStore readOnly(Boolean readOnly) {
    this.readOnly = readOnly;
    return this;
  }

   /**
   * Get readOnly
   * @return readOnly
  **/
  @ApiModelProperty(value = "")
  public Boolean isReadOnly() {
    return readOnly;
  }

  public void setReadOnly(Boolean readOnly) {
    this.readOnly = readOnly;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1ArtifactsStore v1ArtifactsStore = (V1ArtifactsStore) o;
    return Objects.equals(this.uuid, v1ArtifactsStore.uuid) &&
        Objects.equals(this.name, v1ArtifactsStore.name) &&
        Objects.equals(this.description, v1ArtifactsStore.description) &&
        Objects.equals(this.readme, v1ArtifactsStore.readme) &&
        Objects.equals(this.tags, v1ArtifactsStore.tags) &&
        Objects.equals(this.createdAt, v1ArtifactsStore.createdAt) &&
        Objects.equals(this.updatedAt, v1ArtifactsStore.updatedAt) &&
        Objects.equals(this.frozen, v1ArtifactsStore.frozen) &&
        Objects.equals(this.disabled, v1ArtifactsStore.disabled) &&
        Objects.equals(this.deleted, v1ArtifactsStore.deleted) &&
        Objects.equals(this.k8sSecret, v1ArtifactsStore.k8sSecret) &&
        Objects.equals(this.type, v1ArtifactsStore.type) &&
        Objects.equals(this.mountPath, v1ArtifactsStore.mountPath) &&
        Objects.equals(this.hostPath, v1ArtifactsStore.hostPath) &&
        Objects.equals(this.volumeClaim, v1ArtifactsStore.volumeClaim) &&
        Objects.equals(this.bucket, v1ArtifactsStore.bucket) &&
        Objects.equals(this.readOnly, v1ArtifactsStore.readOnly);
  }

  @Override
  public int hashCode() {
    return Objects.hash(uuid, name, description, readme, tags, createdAt, updatedAt, frozen, disabled, deleted, k8sSecret, type, mountPath, hostPath, volumeClaim, bucket, readOnly);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1ArtifactsStore {\n");
    
    sb.append("    uuid: ").append(toIndentedString(uuid)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    readme: ").append(toIndentedString(readme)).append("\n");
    sb.append("    tags: ").append(toIndentedString(tags)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
    sb.append("    frozen: ").append(toIndentedString(frozen)).append("\n");
    sb.append("    disabled: ").append(toIndentedString(disabled)).append("\n");
    sb.append("    deleted: ").append(toIndentedString(deleted)).append("\n");
    sb.append("    k8sSecret: ").append(toIndentedString(k8sSecret)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    mountPath: ").append(toIndentedString(mountPath)).append("\n");
    sb.append("    hostPath: ").append(toIndentedString(hostPath)).append("\n");
    sb.append("    volumeClaim: ").append(toIndentedString(volumeClaim)).append("\n");
    sb.append("    bucket: ").append(toIndentedString(bucket)).append("\n");
    sb.append("    readOnly: ").append(toIndentedString(readOnly)).append("\n");
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

