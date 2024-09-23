#version 330

layout(location = 0) in vec3 vertexPosition;
layout(location = 1) in vec2 vertexTexCoord;
layout(location = 2) in vec3 vertexNormal;
layout(location = 3) in vec4 vertexColor;

out vec2 fragTexCoord;
out vec4 fragColor;
out vec3 fragNormal;
out float time;

uniform mat4 matModel;
uniform mat4 matView;
uniform mat4 matProjection;

uniform float u_time;

void main() {
    gl_Position = matProjection * matView * matModel * vec4(vertexPosition, 1.0);

    float pulse = 0.05 * abs(sin(u_time * 3));

    fragTexCoord = vec2(vertexTexCoord.x, vertexTexCoord.y + pulse);
    fragColor = vertexColor;
    fragNormal = vertexNormal;
    time = u_time;
}
