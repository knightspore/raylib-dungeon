#version 330

in vec2 fragTexCoord;
in vec4 fragColor;

out vec4 finalColor;

uniform sampler2D tex;
uniform sampler2D u_normal;
uniform vec2 u_resolution;
uniform vec2 u_lightPos;
uniform float u_zoom;

void main() {
    // Setup

    // Albedo, Normal, Specular, Position
    vec4 textureColor = texture(tex, fragTexCoord);
    vec4 normal = texture(u_normal, fragTexCoord);
    float specular = pow(dot(textureColor.rgb, vec3(0.299, 0.587, 0.114)), 2.0);
    vec2 position = (fragTexCoord * u_resolution); // screen space

    // Lighting

    // Create normalized position from uniform, flipping y as per OpenGL coordinate system
    vec2 lightPos = vec2(u_lightPos.x / u_resolution.x, 1.0 - u_lightPos.y / u_resolution.y);

    // Debug

    // Draw red circle around light pos
    bool withinRad = distance(fragTexCoord, lightPos) < 0.1 * u_zoom; // 0.1 is the radius
    if (withinRad) {
        textureColor.rg = fragTexCoord.xy;
    }

    finalColor = textureColor;
}
