#version 330

layout(location = 0) in vec3 vertexPosition;
layout(location = 1) in vec2 vertexTexCoord;
layout(location = 2) in vec3 vertexNormal;
layout(location = 3) in vec4 vertexColor;

out vec2 fragTexCoord;
out vec4 fragColor;
out vec3 fragNormal;

uniform mat4 matModel;
uniform mat4 matView;
uniform mat4 matProjection;

void main() {
    gl_Position = matProjection * matView * matModel * vec4(vertexPosition, 1.0);
    fragTexCoord = vertexTexCoord;
    fragColor = vertexColor;
    fragNormal = vertexNormal;
}
