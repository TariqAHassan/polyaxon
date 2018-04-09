# Generated by Django 2.0.3 on 2018-04-09 09:49

from django.conf import settings
import django.core.validators
from django.db import migrations, models
import django.db.models.deletion
import libs.blacklist
import re
import uuid


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ('plugins', '0001_initial'),
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
    ]

    operations = [
        migrations.CreateModel(
            name='Project',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('description', models.TextField(blank=True, null=True)),
                ('created_at', models.DateTimeField(auto_now_add=True, db_index=True)),
                ('updated_at', models.DateTimeField(auto_now=True)),
                ('uuid', models.UUIDField(default=uuid.uuid4, editable=False, unique=True)),
                ('name', models.CharField(max_length=256, validators=[django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'), "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.", 'invalid'), libs.blacklist.validate_blacklist_name])),
                ('is_public', models.BooleanField(default=True, help_text='If project is public or private.')),
                ('has_tensorboard', models.BooleanField(default=False, help_text='If project has a tensorboard.')),
                ('has_notebook', models.BooleanField(default=False, help_text='If project has a notebook.')),
                ('notebook', models.OneToOneField(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, to='plugins.NotebookJob')),
                ('tensorboard', models.OneToOneField(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, to='plugins.TensorboardJob')),
                ('user', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='projects', to=settings.AUTH_USER_MODEL)),
            ],
        ),
        migrations.AlterUniqueTogether(
            name='project',
            unique_together={('user', 'name')},
        ),
    ]
