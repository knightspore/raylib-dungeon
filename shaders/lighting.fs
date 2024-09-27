#version 410 core

#define NUM_LIGHTS 2

in vec2 fragTexCoord;
in vec4 fragColor;

out vec4 finalColor;

uniform sampler2D tex;
uniform sampler2D u_normal;
uniform vec2 u_resolution;
uniform float u_ambient;
uniform vec2 u_lightPos[NUM_LIGHTS];
uniform vec3 u_lightColor[NUM_LIGHTS];

void main() {
    vec4 textureColor = texture(tex, fragTexCoord);
    vec4 normal = texture(u_normal, fragTexCoord);
    vec3 lighting = textureColor.rgb * u_ambient;

    for (int i = 0; i < NUM_LIGHTS; i++) {
        vec3 lightColor = vec3(u_lightColor[i].r / 255.0, u_lightColor[i].g / 255.0, u_lightColor[i].b / 255.0);
        vec2 pos = vec2(u_lightPos[i].x / u_resolution.x, 1.0 - u_lightPos[i].y / u_resolution.y);
        vec2 lightDir2d = pos - fragTexCoord;
        vec2 lightDirNorm = normalize(lightDir2d);
        vec3 lightDir3 = normalize(vec3(lightDirNorm, 200.0));
        float distance = length(lightDir2d);

        float diffuseStrength = max(dot(normal.rgb, lightDir3), u_ambient);
        float attenuation = 1.0 / (1.0 + 5.0 * distance + 4.0 * distance * distance);
        vec3 diffuse = diffuseStrength * attenuation * textureColor.rgb * lightColor;
        vec3 specular = pow(dot(textureColor.rgb, vec3(0.299, 0.587, 0.114)), 2.0) * attenuation * lightColor;

        lighting += diffuse + specular;
    }

    finalColor = vec4(lighting, textureColor.a);
}
