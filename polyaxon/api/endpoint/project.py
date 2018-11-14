from django.shortcuts import get_object_or_404

import access

from access.entities import Entities
from api.endpoint.base import BaseEndpoint
from api.endpoint.owner import OwnerPermission
from db.models.projects import Project


class ProjectPermission(OwnerPermission):
    SCOPE_MAPPING = access.get_scope_mapping_for('Project')

    def has_object_permission(self, request, view, obj):
        if self._check_internal_or_ephemeral(request=request):
            return True

        result = super().has_object_permission(request, view, obj.owner)
        if not result:
            return result

        return access.has_object_permission(
            entity=Entities.PROJECT,
            permission=ProjectPermission,
            request=request,
            view=view,
            obj=obj)


class ProjectResourceListPermission(ProjectPermission):
    SCOPE_MAPPING = access.get_scope_mapping_for('ProjectResourceListPermission')


class ProjectResourcePermission(ProjectPermission):
    SCOPE_MAPPING = access.get_scope_mapping_for('ProjectResource')

    def has_object_permission(self, request, view, obj):
        return super().has_object_permission(request, view, obj.project)


class ProjectEndpoint(BaseEndpoint):
    queryset = Project.objects
    permission_classes = (ProjectPermission,)
    AUDITOR_EVENT_TYPES = None
    CONTEXT_KEYS = ('owner_name', 'project_name')
    CONTEXT_OBJECTS = ('owner', 'project')
    lookup_field = 'name'
    lookup_url_kwarg = 'project_name'

    def enrich_queryset(self, queryset):
        return queryset.filter(owner__name=self.owner_name)

    def _initialize_context(self):
        self.project = self.get_object()
        self.owner = self.project.owner


class ProjectResourceListEndpoint(BaseEndpoint):
    permission_classes = (ProjectResourceListPermission,)
    AUDITOR_EVENT_TYPES = None
    CONTEXT_KEYS = ('owner_name', 'project_name')
    CONTEXT_OBJECTS = ('owner', 'project')

    def enrich_queryset(self, queryset):
        queryset = queryset.filter(project=self.project)
        return super().enrich_queryset(queryset=queryset)

    def _initialize_context(self):
        super()._initialize_context()
        self.project = get_object_or_404(Project,
                                         owner__name=self.owner_name,
                                         name=self.project_name)
        self.owner = self.project.owner

    def _validate_context(self):
        super()._validate_context()
        permission = self.permission_classes[0]()
        if not permission.has_object_permission(request=self.request,
                                                view=self,
                                                obj=self.project):
            self.permission_denied(
                self.request, message=getattr(permission, 'message', None)
            )


class ProjectResourceEndpoint(ProjectResourceListEndpoint):
    permission_classes = (ProjectResourcePermission,)
    AUDITOR_EVENT_TYPES = None
    lookup_field = 'id'

    def _validate_context(self):
        super(ProjectResourceListEndpoint, self)._validate_context()